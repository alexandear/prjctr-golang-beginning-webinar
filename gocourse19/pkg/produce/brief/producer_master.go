package brief

import (
	"context"
	"errors"
	"log"

	"github.com/streadway/amqp"
)

type Master struct {
	exName     string
	connection *amqp.Connection
	done       chan struct{}
	isReady    bool
	ackTag     chan uint64
}

func NewMaster(ctx context.Context, exName, busHost, busUser, busPass string) (*Master, error) {
	session := Master{
		exName: exName,
		done:   make(chan struct{}),
	}

	err := session.connect(ctx, busHost, busUser, busPass)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (s *Master) connect(_ context.Context, busHost, busUser, busPass string) error {
	conn, err := amqp.DialConfig(busHost, amqp.Config{
		SASL: []amqp.Authentication{&amqp.PlainAuth{busUser, busPass}},
	})
	if err != nil {
		return err
	}

	s.connection = conn
	s.isReady = true
	s.done = make(chan struct{})
	log.Println("Producer: CONNECTED")

	return nil
}

func (s *Master) Close() error {
	if !s.isReady {
		return errors.New("producer: connection not ready while closing")
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
