package tasks

import (
	"context"
	"github.com/samber/do"
	"gocourse22/pkg/scheduler"
)

func CalculateEmployysBonuses(_ *do.Injector) *CalculateEmployys {
	return &CalculateEmployys{}
}

type CalculateEmployys struct {
}

func (r *CalculateEmployys) TimeType() scheduler.TimeType {
	return scheduler.Every
}

func (r *CalculateEmployys) Expression() string {
	return `1m`
}

func (r *CalculateEmployys) Operation(_ context.Context, inj *do.Injector) func() {
	return func() {
	}
}
