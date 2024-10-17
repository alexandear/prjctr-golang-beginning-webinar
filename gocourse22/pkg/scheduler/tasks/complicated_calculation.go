package tasks

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/samber/do"
	"gocourse22/pkg/scheduler"
	"golang.org/x/sync/errgroup"
	"log"
)

func NewComplicatedCalculation(_ *do.Injector) *ComplicatedCalculation {
	return &ComplicatedCalculation{}
}

type ComplicatedCalculation struct {
}

func (r *ComplicatedCalculation) TimeType() scheduler.TimeType {
	return scheduler.Every
}

func (r *ComplicatedCalculation) Expression() string {
	return `1m`
}

func (r *ComplicatedCalculation) Operation(ctx context.Context, inj *do.Injector) func() {
	db := do.MustInvokeNamed[*pgxpool.Pool](inj, "postgres")

	return func() {
		// Ініціалізація errgroup
		g, ctx := errgroup.WithContext(ctx)

		// Симуляція операцій для різних сутностей
		g.Go(func() error {
			return fetchDoctors(ctx, db)
		})
		g.Go(func() error {
			return fetchPatients(ctx, db)
		})
		g.Go(func() error {
			return fetchVisits(ctx, db)
		})
		g.Go(func() error {
			return fetchPrescriptions(ctx, db)
		})

		// Чекаємо завершення всіх горутин
		if err := g.Wait(); err != nil {
			log.Fatalf("encountered error: %v", err)
		}
		fmt.Println("All operations completed successfully")
	}
}

func fetchDoctors(ctx context.Context, db *pgxpool.Pool) error {
	// Тут міг би бути ваш код для зчитування даних про лікарів
	fmt.Println("Fetching doctors...")
	return nil // Замініть на реальну реалізацію
}

func fetchPatients(ctx context.Context, db *pgxpool.Pool) error {
	// Тут міг би бути ваш код для зчитування даних про пацієнтів
	fmt.Println("Fetching patients...")
	return nil // Замініть на реальну реалізацію
}

func fetchVisits(ctx context.Context, db *pgxpool.Pool) error {
	// Тут міг би бути ваш код для зчитування даних про візити
	fmt.Println("Fetching visits...")
	return nil // Замініть на реальну реалізацію
}

func fetchPrescriptions(ctx context.Context, db *pgxpool.Pool) error {
	// Тут міг би бути ваш код для зчитування даних про призначення ліків
	fmt.Println("Fetching prescriptions...")
	return nil // Замініть на реальну реалізацію
}
