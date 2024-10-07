package produce

import (
	"context"
	p "gocourse19/pkg/produce/full"
	"log"
	"time"
)

type Pool struct {
	master *p.Master
	slaves []*p.Slave
}

func NewPool(ctx context.Context, threads int, exName, busHost, busUser, busPass string) *Pool {
	master, err := p.NewMaster(ctx, exName, busHost, busUser, busPass)
	if err != nil {
		log.Fatalln(err)
	}
	res := &Pool{master: master}

	for i := 0; i < threads; i++ {
		slave, err := p.NewSlave(ctx, res.master)
		if err != nil {
			log.Fatalln(err)
		}
		res.slaves = append(res.slaves, slave)
	}

	return res
}

func (s *Pool) Producers() []*p.Slave {
	return s.slaves
}

func (s *Pool) Close(_ context.Context) {
	log.Println("Producer: close slaves")
	for _, slave := range s.slaves {
		slave.Complete()
		err := slave.Close()
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Producer: slaves were closed, close master")
	s.master.Complete()
	time.Sleep(3 * time.Second)
	err := s.master.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println("Producer: master closed")
}
