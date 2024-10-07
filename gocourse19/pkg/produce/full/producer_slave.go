package full

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	channelReconnectDelay = 5 * time.Second
	reInitDelay           = 5 * time.Second
	resendDelay           = 15 * time.Second
	confirmRetries        = 9
	pushRetries           = 3
)

type Slave struct {
	master          *Master
	channel         *amqp.Channel
	done            chan struct{}
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	notifyFlow      chan bool
	IsReady         bool
	tm              *time.Ticker
}

var (
	errShutdown     = errors.New("-- Producer session shut down")
	errConnNotReady = errors.New("Producer: connection not ready")
)

func NewSlave(ctx context.Context, master *Master) (*Slave, error) {
	session := Slave{
		master: master,
		done:   make(chan struct{}),
		tm:     time.NewTicker(resendDelay),
	}
	go session.handleReInit(ctx)

	return &session, nil
}

func (s *Slave) handleReInit(ctx context.Context) bool {
	for {
		s.IsReady = false

		err := s.init(ctx)
		if err != nil {
			log.Println("Producer: failed to initialize channel (%s). Retrying...", err)

			select {
			case <-s.master.done:
				return true
			case <-s.done:
				return true
			case <-time.After(reInitDelay):
			}
			continue
		}

		select {
		case <-s.master.done:
			return true
		case <-s.done:
			return true
		case <-s.notifyChanClose:
			log.Println("Producer: channel closed. Re-init...")
		}
	}
}

func (s *Slave) init(ctx context.Context) error {
	for {
		if s.master.connection == nil || s.master.connection.IsClosed() {
			log.Println("Producer: connection not ready. Waiting...")
			time.Sleep(channelReconnectDelay)
		} else {
			break
		}
	}

	ch, err := s.master.connection.Channel()
	if err != nil {
		return err
	}

	err = ch.Confirm(false)
	if err != nil {
		return err
	}

	err = s.declarationAndBinding(ctx, ch)
	if err != nil {
		return err
	}

	s.changeChannel(ctx, ch)
	s.IsReady = true
	s.done = make(chan struct{})
	log.Println("Producer: SETUP")

	return nil
}

func (s *Slave) declarationAndBinding(_ context.Context, ch *amqp.Channel) (err error) {
	queues := []string{`q1`, `q2`, `q3`}
	queuesEntities := map[string][]string{
		`q1`: {`product`, `brand`},
		`q2`: {`category`},
		`q3`: {`product`, `attribute`},
	}

	for _, qName := range queues {
		_, err = ch.QueueDeclare(qName, true, false, false, false, nil)
		if err != nil {
			return
		}
	}

	for qName, entities := range queuesEntities {
		for _, entity := range entities {
			err = ch.QueueBind(qName, fmt.Sprintf("key-%s", entity), s.master.exName, false, nil)
			if err != nil {
				return
			}
		}
	}

	return
}

func (s *Slave) changeChannel(ctx context.Context, channel *amqp.Channel) {
	s.channel = channel
	s.notifyChanClose = make(chan *amqp.Error, 1)
	s.notifyConfirm = make(chan amqp.Confirmation, 1)
	s.channel.NotifyClose(s.notifyChanClose)
	s.channel.NotifyPublish(s.notifyConfirm)

	// research block, is this notification will be flashed
	s.notifyFlow = make(chan bool, 1)
	s.channel.NotifyFlow(s.notifyFlow)

	go s.listenFlow(ctx)
}

func (s *Slave) listenFlow(_ context.Context) {
	for {
		select {
		case res, ok := <-s.notifyFlow:
			log.Println("Producer: receive notifyFlow = %v, is closed = %v", res, ok)
			if !ok {
				return
			}
		}
	}
}

func (s *Slave) Push(_ context.Context, rk string, body []byte) error {
	tm := time.NewTicker(resendDelay)
	defer tm.Stop()

	retries := 0
	for {
		if !s.IsReady {
			if retries > pushRetries {
				return errors.New("Producer: failed to push")
			} else {
				log.Println("Producer: failed to push. Retrying...")
				retries++
				time.Sleep(channelReconnectDelay)
			}
		} else {
			break
		}
	}

	retries = 0
	for {
		if !s.IsReady {
			return errConnNotReady
		}

		err := s.UnsafePush(rk, body)
		if err != nil {
			log.Println("Producer: Push failed: %s. (%s) Retrying...", err, rk)
			select {
			case <-s.master.done:
				log.Println("receive done signal from master %s", rk)
				return errShutdown
			case <-s.done:
				log.Println("receive done signal %s", rk)
				return errShutdown
			case <-tm.C:
			}
			continue
		}

		for {
			if !s.IsReady {
				return errConnNotReady
			}
			select {
			case confirm := <-s.notifyConfirm:
				if confirm.Ack {
					log.Println("Producer: published successfully into %s", rk)
					return nil
				} else {
					log.Println("producxer_slave, NOT Acked to %s", rk)
				}
			case <-s.master.done:
				log.Println("receive done signal from master to %s", rk)
				return nil
			case <-s.done:
				log.Println("receive done signal to %s", rk)
				return nil
			case <-tm.C:
				log.Println("producer_slave, relisten to %s", rk)
			}
			if s.master.connection.IsClosed() {
				return errConnNotReady
			}
			if retries > confirmRetries {
				return fmt.Errorf("Producer: failed to confirm to %s", rk)
			} else {
				retries++
				log.Println("Producer: failed to confirm. Retrying... to %s", rk)
			}
		}
	}
}

func (s *Slave) UnsafePush(rk string, body []byte) error {
	if !s.IsReady {
		return errors.New(fmt.Sprintf("Producer: connection not ready"))
	}

	return s.channel.Publish(
		s.master.exName,
		rk,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Body:         body,
			Priority:     5,
		},
	)
}

func (s *Slave) Close() error {
	if !s.IsReady {
		return errors.New(fmt.Sprintf("Producer: channel not ready while closing"))
	}
	err := s.channel.Close()
	if err != nil {
		return err
	}
	s.IsReady = false

	return nil
}

func (s *Slave) Complete() {
	close(s.done)
}
