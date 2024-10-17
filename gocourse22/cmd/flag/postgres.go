package flag

import "github.com/urfave/cli/v2"

const (
	PostgresHost   = "postgres-host"
	PostgresPort   = "postgres-port"
	PostgresUser   = "postgres-user"
	PostgresPass   = "postgres-pass"
	PostgresDBName = "postgres-db-name"
)

func PostgresHostFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "database connection",
		Name:        PostgresHost,
		Usage:       "Postgres host",
		EnvVars:     []string{"POSTGRES_HOST"},
		DefaultText: "localhost",
		Value:       "localhost",
	}
}

func PostgresPortFlag() cli.Flag {
	return &cli.IntFlag{
		Category:    "database connection",
		Name:        PostgresPort,
		Usage:       "Postgres port",
		EnvVars:     []string{"POSTGRES_PORT"},
		Value:       5432,
		DefaultText: "5432",
	}
}

func PostgresUserFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "database connection",
		Name:        PostgresUser,
		Usage:       "Postgres user",
		EnvVars:     []string{"POSTGRES_USER"},
		Value:       "postgres",
		DefaultText: "postgres",
	}
}

func PostgresPasswordFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "database connection",
		Name:        PostgresPass,
		Usage:       "Postgres password",
		EnvVars:     []string{"POSTGRES_PASS"},
		Value:       "123456",
		DefaultText: "123456",
	}
}

func PostgresDBNameFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "database connection",
		Name:        PostgresDBName,
		Usage:       "Postgres database name",
		EnvVars:     []string{"POSTGRES_DB_NAME"},
		Value:       "gocourse22",
		DefaultText: "gocourse22",
	}
}
