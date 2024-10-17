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
	"prjctr.com/gocourse22/internal/providers"
	"prjctr.com/gocourse22/pkg/scheduler"
	"prjctr.com/gocourse22/pkg/scheduler/tasks"
)

// Worker define the run command
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
			ctx, cancelFunc := signal.NotifyContext(c.Context, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
			defer cancelFunc()

			// inject the signal notify context
			do.ProvideValue(injector, ctx)

			// needed to use flags provided by the cmd.Run command
			c.Context = ctx
			do.OverrideValue(injector, c)

			providers.ProvideConnection(injector)
			do.Provide(injector, providers.ProvideScheduler)

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
					tasks.CalculateEmployysBonuses(injector),
				); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						log.Fatal(err)
					}
				}
				log.Println("Scheduler has been stopped")
			}()

			stopWg.Wait()
			return injector.Shutdown()
		},
	}
}
