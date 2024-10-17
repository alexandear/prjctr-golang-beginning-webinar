package provider

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"

	"prjctr.com/gocourse22/cmd/flag"
	"prjctr.com/gocourse22/internal/db"
)

func Connection(injector *do.Injector) {
	do.ProvideNamed(injector, "postgres", PostgresConnection)
}

func PostgresConnection(injector *do.Injector) (*pgxpool.Pool, error) {
	c := do.MustInvoke[*cli.Context](injector)

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.String(flag.PostgresHost),
		c.Int(flag.PostgresPort),
		c.String(flag.PostgresUser),
		c.String(flag.PostgresPass),
		c.String(flag.PostgresDBName),
	)

	return db.NewConnectionPool(c.Context, dsn)
}
