package cmd

import (
	"fmt"
	"log"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"

	"prjctr.com/gocourse22/cmd/flag"
	"prjctr.com/gocourse22/internal/provider"
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

// MigrationApply applies the migrations via the atlas binary.
// Read the documentation at https://atlasgo.io/getting-started/#installation to learn how to install it.
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

			provider.Connection(injector)

			wd, _ := os.Getwd()
			migrationsSource := fmt.Sprintf("file://%s/migrations", wd)

			client, clientErr := atlasexec.NewClient(".", c.String(flag.AtlasBin))
			if clientErr != nil {
				return fmt.Errorf("initialize atlas client: %w", clientErr)
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
				Schema:      []string{"public"},
				To:          migrationsSource,
				URL:         dsnURL,
				AutoApprove: true,
			}

			res, resErr := client.SchemaApply(c.Context, &applyParams)
			if resErr != nil {
				return fmt.Errorf("apply schema: %w", resErr)
			}

			var appliedCount int
			if res != nil {
				appliedCount = len(res.Changes.Applied)
			}
			log.Printf("Applied %d migrations", appliedCount)

			return nil
		},
	}
}
