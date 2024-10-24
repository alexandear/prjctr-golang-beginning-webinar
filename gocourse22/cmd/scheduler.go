package cmd

import (
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"

	"github.com/samber/do"
	"github.com/urfave/cli/v2"

	"prjctr.com/gocourse22/cmd/flag"
	"prjctr.com/gocourse22/internal/provider"
	"prjctr.com/gocourse22/pkg/scheduler"
	"prjctr.com/gocourse22/pkg/scheduler/tasks"
)

// Worker define the run command.
func Worker() *cli.Command {
	return &cli.Command{
		Name:  "scheduler",
		Usage: "The Worker",
		Flags: []cli.Flag{
			// db
			flag.PostgresHostFlag(),
			flag.PostgresPortFlag(),
			flag.PostgresUserFlag(),
			flag.PostgresPasswordFlag(),
			flag.PostgresDBNameFlag(),

			flag.InstanceEnvFlag(),
			flag.InstanceIDFlag(),
			flag.VarDirectoryFlag(),
		},
		Action: func(c *cli.Context) error {
			// create injector
			injector := do.DefaultInjector

			// listen to os interrupt signals and close the context
			ctx, cancel := signal.NotifyContext(c.Context, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
			defer cancel()

			// inject the signal notify context
			do.ProvideValue(injector, ctx)

			// needed to use flags provided by the cmd.Run command
			c.Context = ctx
			do.OverrideValue(injector, c)

			provider.Connection(injector)
			do.Provide(injector, provider.Scheduler)

			stopWg := sync.WaitGroup{}
			stopWg.Add(1)
			// start the scheduler service
			go func() {
				defer stopWg.Done()

				tasksScheduler := do.MustInvoke[*scheduler.Scheduler](injector)
				go func() {
					<-ctx.Done()
					tasksScheduler.Shutdown()
				}()

				if err := tasksScheduler.Manage(ctx,
					tasks.NewComplicatedCalculation(injector),
					tasks.GenerateEmployeeBonuses(injector),
				); !errors.Is(err, http.ErrServerClosed) {
					log.Printf("Failed to manage tasks: %v\n", err)
					return
				}
				log.Println("Scheduler has been stopped")
			}()

			stopWg.Wait()
			return injector.Shutdown()
		},
	}
}
