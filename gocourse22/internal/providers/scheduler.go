package providers

import (
	"github.com/samber/do"

	"prjctr.com/gocourse22/pkg/scheduler"
)

func ProvideScheduler(i *do.Injector) (*scheduler.Scheduler, error) {
	return scheduler.NewScheduler(i), nil
}
