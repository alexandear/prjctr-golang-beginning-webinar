package tasks

import (
	"context"

	"github.com/samber/do"

	"prjctr.com/gocourse22/pkg/scheduler"
)

func GenerateEmployeeBonuses(_ *do.Injector) *CalculateEmployees {
	return &CalculateEmployees{}
}

type CalculateEmployees struct{}

func (r *CalculateEmployees) TimeType() scheduler.TimeType {
	return scheduler.Every
}

func (r *CalculateEmployees) Expression() string {
	return `1m`
}

func (r *CalculateEmployees) Operation(_ context.Context, inj *do.Injector) func() {
	return func() {
	}
}
