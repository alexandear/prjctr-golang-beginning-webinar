package full

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

const (
	ChannelReconnectDelay = 5 * time.Second
	reInitDelay           = 2 * time.Second
)

var errNotConnected = errors.New("consumer: not connected to the server")

type Slave struct {
	master          *Master
	channel         *amqp.Channel
	done            chan struct{}
	closed          chan struct{}
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	delivery        <-chan amqp.Delivery
	isReady         bool
	IsDeliveryReady bool
	chMux           *sync.Mutex
	prefetchCount   int
}

func NewSlave(ctx context.Context, master *Master, prefetchCount int) (*Slave, error) {
	session := Slave{
		master:        master,
		done:          make(chan struct{}),
		chMux:         &sync.Mutex{},
		prefetchCount: prefetchCount,
	}
	go session.handleReInit(ctx)

	return &session, nil
}

func (s *Slave) handleReInit(ctx context.Context) bool {
	for {
		s.isReady = false
		s.IsDeliveryReady = false

		err := s.init(ctx)
		if err != nil {
			log.Printf("Consumer: failed to initialize channel (%s). Retrying...\n", err)

			select {
			case <-s.done:
				return true
			case <-time.After(reInitDelay):
			}
			continue
		}

		err = s.InitStream(ctx)
		if err != nil {
			log.Printf("Consumer: stream not inited: %s\n", err)
		}

		select {
		case <-s.done:
			return true
		case <-s.notifyConnClose:
			s.closed <- struct{}{}
			log.Println("Consumer: connection closed. Reconnecting...")
		case <-s.notifyChanClose:
			s.closed <- struct{}{}
			log.Println("Consumer: channel closed. Re-running init...")
		}
	}
}

func (s *Slave) init(ctx context.Context) error {
	for {
		if s.master.connection == nil || s.master.connection.IsClosed() {
			log.Println("Consumer: connection not ready. Waiting...")
			time.Sleep(ChannelReconnectDelay)
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

	err = ch.Qos(s.prefetchCount, 0, false)
	if err != nil {
		return err
	}

	s.changeChannel(ch)
	s.isReady = true
	s.closed = make(chan struct{})
	log.Println("Consumer: SETUP")

	return nil
}

func (s *Slave) changeChannel(channel *amqp.Channel) {
	s.channel = channel
	s.notifyChanClose = make(chan *amqp.Error, 1)
	s.notifyConfirm = make(chan amqp.Confirmation, 1)
	s.channel.NotifyClose(s.notifyChanClose)
	s.channel.NotifyPublish(s.notifyConfirm)
}

func (s *Slave) InitStream(_ context.Context) (err error) {
	if !s.isReady {
		return errNotConnected
	}

	log.Printf("consume queue %s/%s\n", s.master.connection.Config.Vhost, s.master.qName)
	s.chMux.Lock()
	s.delivery, err = s.channel.Consume(
		s.master.qName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	s.chMux.Unlock()

	if err == nil {
		log.Println("Consumer: Stream SETUP")
		s.IsDeliveryReady = true
	}

	return
}

func (s *Slave) GetStream() <-chan amqp.Delivery {
	s.chMux.Lock()
	d := s.delivery
	s.chMux.Unlock()

	return d
}

func (s *Slave) Closed() <-chan struct{} {
	return s.closed
}

func (s *Slave) Close() error {
	if !s.isReady {
		return errors.New("consumer: channel not ready while closing")
	}
	err := s.channel.Close()
	if err != nil {
		return err
	}
	s.isReady = false
	s.IsDeliveryReady = false

	return nil
}

func (s *Slave) Complete() {
	s.done <- struct{}{}
}
