package cmd

import (
	"fmt"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"

	"github.com/samber/do"
	"github.com/urfave/cli/v2"
	"gocourse22/cmd/flag"
	common "gocourse22/internal/providers"
)

func Migrate() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Apply migrations",
		Flags: []cli.Flag{},
		Subcommands: []*cli.Command{
			MigrateApply(),
		},
		Action: nil,
	}
}

func MigrateApply() *cli.Command {
	return &cli.Command{
		Name:  "apply",
		Usage: "Apply database migrations",
		Flags: []cli.Flag{
			flag.AtlasBinFlag(),

			flag.PostgresHostFlag(),
			flag.PostgresPortFlag(),
			flag.PostgresUserFlag(),
			flag.PostgresPasswordFlag(),
			flag.PostgresDBNameFlag(),
		},
		Action: func(c *cli.Context) (err error) {
			// create injector
			injector := do.DefaultInjector
			defer func() {
				err = injector.Shutdown()
			}()

			do.OverrideValue(injector, c)

			common.ProvideConnection(injector)

			wd, _ := os.Getwd()
			migrationsSource := fmt.Sprintf("file://%s/migrations", wd)

			client, clientErr := atlasexec.NewClient(".", c.String(flag.AtlasBin))
			if clientErr != nil {
				log.Fatalf("failed to initialize atlas client: %v", clientErr)
				return clientErr
			}

			dsnURL := fmt.Sprintf(
				"postgres://%s:%s@%s:%d/%s?search_path=public&sslmode=disable",
				c.String(flag.PostgresUser),
				c.String(flag.PostgresPass),
				c.String(flag.PostgresHost),
				c.Int(flag.PostgresPort),
				c.String(flag.PostgresDBName),
			)

			applyParams := atlasexec.SchemaApplyParams{
				Schema: []string{"public"},
				To:     migrationsSource,
				URL:    dsnURL,
			}

			res, resErr := client.SchemaApply(c.Context, &applyParams)
			if resErr != nil {
				log.Fatalf("Failed to apply schema due to error: %v", resErr)
				return resErr
			}

			if res == nil {
				log.Println("Applied 0 migrations")
			} else {
				log.Printf("Applied %d migrations", len(res.Changes.Applied))
			}

			return
		},
	}
}
