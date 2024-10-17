package scheduler

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/samber/do"
)

type Task interface {
	TimeType() TimeType
	Expression() string
	Operation(ctx context.Context, inj *do.Injector) func()
}

type TimeType uint8

const (
	Every TimeType = iota
	Cron
)

func (tt TimeType) String() string {
	switch tt {
	case Every:
		return "every"
	case Cron:
		return "cron"
	default:
		return "unknown"
	}
}

func NewScheduler(inj *do.Injector) *Scheduler {
	return &Scheduler{inj: inj}
}

type Scheduler struct {
	inj        *do.Injector
	_scheduler *gocron.Scheduler
}

func (r *Scheduler) Manage(ctx context.Context, tasks ...Task) error {
	s := gocron.NewScheduler(time.UTC)

	for i := range tasks {
		switch tasks[i].TimeType() {
		case Every:
			_, err := s.Every(tasks[i].Expression()).Do(tasks[i].Operation(ctx, r.inj))
			if err != nil {
				return err
			}
		case Cron:
			_, err := s.Cron(tasks[i].Expression()).Do(tasks[i].Operation(ctx, r.inj))
			if err != nil {
				return err
			}
		}
	}

	s.StartAsync()

	<-ctx.Done()

	r._scheduler = s

	return nil
}

func (r *Scheduler) Shutdown() {
	if r._scheduler != nil {
		r._scheduler.Stop()
	}
}
