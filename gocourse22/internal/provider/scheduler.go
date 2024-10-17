package provider

import (
	"github.com/samber/do"

	"prjctr.com/gocourse22/pkg/scheduler"
)

func Scheduler(injector *do.Injector) (*scheduler.Scheduler, error) {
	return scheduler.NewScheduler(injector), nil
}
