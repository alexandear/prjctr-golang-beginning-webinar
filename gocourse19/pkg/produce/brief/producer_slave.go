package brief

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
	resendDelay           = 15 * time.Second
	pushRetries           = 3
)

type Slave struct {
	master  *Master
	channel *amqp.Channel
	done    chan struct{}
	IsReady bool
	tm      *time.Ticker
}

func NewSlave(ctx context.Context, master *Master) (*Slave, error) {
	session := Slave{
		master: master,
		done:   make(chan struct{}),
		tm:     time.NewTicker(resendDelay),
	}

	err := session.init(ctx)
	if err != nil {
		return nil, err
	}

	return &session, nil
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

	s.channel = ch
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

func (s *Slave) Push(rk string, body []byte) error {
	if !s.IsReady {
		return errors.New("producer: connection not ready")
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
		return errors.New("Producer: channel not ready while closing")
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
