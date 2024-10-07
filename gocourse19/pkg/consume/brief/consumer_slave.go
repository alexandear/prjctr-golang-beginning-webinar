package brief

import (
	"context"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

const (
	ChannelReconnectDelay = 5 * time.Second
)

var (
	errNotConnected = errors.New("consumer: not connected to the server")
)

type Slave struct {
	master          *Master
	channel         *amqp.Channel
	done            chan struct{}
	closed          chan struct{}
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

	err := session.init(ctx)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Slave) init(_ context.Context) error {
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

	s.channel = ch
	s.isReady = true
	s.closed = make(chan struct{})
	log.Println("Consumer: SETUP")

	return nil
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
		return errors.New(fmt.Sprintf("Consumer: channel not ready while closing"))
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
