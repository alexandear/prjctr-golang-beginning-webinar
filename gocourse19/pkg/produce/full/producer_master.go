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
	connectionReconnectDelay = 5 * time.Second
)

type Master struct {
	exName          string
	connection      *amqp.Connection
	done            chan struct{}
	notifyConnClose chan *amqp.Error
	isReady         bool
	ackTag          chan uint64
}

func NewMaster(ctx context.Context, exName, busHost, busUser, busPass string) (*Master, error) {
	session := Master{
		exName: exName,
		done:   make(chan struct{}),
	}

	go session.handleReconnect(ctx, busHost, busUser, busPass)

	return &session, nil
}

func (s *Master) handleReconnect(ctx context.Context, busHost, busUser, busPass string) {
	for {
		s.isReady = false

		err := s.connect(ctx, busHost, busUser, busPass)
		if err != nil {
			select {
			case <-s.done:
				return
			case <-time.After(connectionReconnectDelay):
			}
			continue
		}

		select {
		case <-s.done:
			return
		case <-s.notifyConnClose:
			log.Println("Producer: connection closed. Reconnecting...")
		}
	}
}

func (s *Master) connect(_ context.Context, busHost, busUser, busPass string) error {
	conn, err := amqp.DialConfig(busHost, amqp.Config{
		SASL: []amqp.Authentication{&amqp.PlainAuth{busUser, busPass}},
	})

	if err != nil {
		return err
	}

	s.changeConnection(conn)
	s.isReady = true
	s.done = make(chan struct{})
	log.Println("Producer: CONNECTED")

	return nil
}

func (s *Master) changeConnection(connection *amqp.Connection) {
	s.connection = connection
	s.notifyConnClose = make(chan *amqp.Error, 1)
	s.connection.NotifyClose(s.notifyConnClose)
}

func (s *Master) Close() error {
	if !s.isReady {
		return errors.New(fmt.Sprintf("Producer: connection not ready while closing"))
	}
	err := s.connection.Close()
	if err != nil {
		return err
	}
	s.isReady = false

	return nil
}

func (s *Master) Complete() {
	close(s.done)
}
