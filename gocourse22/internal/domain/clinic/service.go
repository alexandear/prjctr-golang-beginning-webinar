package clinic

import (
	"context"
	"sync"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/do"

	"prjctr.com/gocourse22/pkg/extend"
)

const (
	visits int = 110
	weeks  int = 3
)

func ProvideService(inj *do.Injector) (*Service, error) {
	return NewService(
		do.MustInvokeNamed[*pgxpool.Pool](inj, "postgres"),
	), nil
}

func NewService(conn *pgxpool.Pool) *Service {
	return &Service{conn}
}

type Service struct {
	conn *pgxpool.Pool
}

type GroupedVisits struct {
	Week  int
	Count int
}

func (s *Service) GroupPatientsVisits() []GroupedVisits {
	visitsCount := make(map[int]int)
	workers := 7

	chanStrategy(workers, visitsCount)

	var res []GroupedVisits
	for week, count := range visitsCount {
		res = append(res, GroupedVisits{week + 1, count})
	}

	return res
}

func (s *Service) GetAll(ctx context.Context) ([]Clinic, error) {
	var res []Clinic
	if err := pgxscan.Select(ctx, s.conn, &res, "SELECT * FROM "+tableName); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *Service) DeleteClinic() error {
	return extend.NewFormattedError(1, "Clinic deletion is impossible", nil)
}

func muxStrategy(workers int, visitsResult map[int]int) { //nolint:unused
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := range workers {
		wg.Add(1)
		go func(week int) {
			defer wg.Done()
			for range visits / weeks {
				mutex.Lock()
				visitsResult[week]++
				mutex.Unlock()
			}
		}(i % weeks)
	}

	wg.Wait()
}

func chanStrategy(workers int, visitsResult map[int]int) {
	results := make(chan map[int]int)

	for i := range workers {
		go func(week int) {
			result := make(map[int]int)
			for range visits / weeks {
				result[week]++
			}
			results <- result
		}(i % weeks)
	}

	for range workers {
		for week, count := range <-results {
			visitsResult[week] += count
		}
	}
}
