package consume

import (
	"context"
	c "gocourse19/pkg/consume/full"
	"log"
	"time"
)

type Pool struct {
	master *c.Master
	slaves []*c.Slave
}

func NewPool(ctx context.Context, threads, prefetchCount int, qName, busHost, busUser, busPass string) *Pool {
	master, err := c.NewMaster(ctx, qName, busHost, busUser, busPass)
	if err != nil {
		log.Fatalln(err)
	}
	res := &Pool{master: master}

	for i := 0; i < threads; i++ {
		slave, err := c.NewSlave(ctx, res.master, prefetchCount)
		if err != nil {
			log.Fatalln(err)
		}
		res.slaves = append(res.slaves, slave)
	}

	return res
}

func (s *Pool) Consumers() []*c.Slave {
	return s.slaves
}

func (s *Pool) Close(ctx context.Context) {
	for _, slave := range s.slaves {
		slave.Complete()
		err := slave.Close()
		if err != nil {
			log.Println(err)
		}
	}
	s.master.Complete()
	err := s.master.Close()
	if err != nil {
		log.Println(err)
	}
}

func (s *Pool) ChannelReconnectDelay() time.Duration {
	return c.ChannelReconnectDelay
}
