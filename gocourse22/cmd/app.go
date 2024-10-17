package cmd

import (
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/samber/do"
	"github.com/urfave/cli/v2"

	"prjctr.com/gocourse22/cmd/flag"
	"prjctr.com/gocourse22/internal/domain/clinic"
	appHttp "prjctr.com/gocourse22/internal/interface/http"
	"prjctr.com/gocourse22/internal/provider"
	"prjctr.com/gocourse22/pkg/extend"
)

// Run define the run command.
func Run() *cli.Command {
	return &cli.Command{
		Name:  "app",
		Usage: "Run the application",
		Flags: []cli.Flag{
			flag.HTTPServerAddressFlag(),
			flag.HTTPReadTimeoutFlag(),
			flag.HTTPShutdownTimeoutFlag(),

			// db
			flag.PostgresHostFlag(),
			flag.PostgresPortFlag(),
			flag.PostgresUserFlag(),
			flag.PostgresPasswordFlag(),
			flag.PostgresDBNameFlag(),
		},
		Action: func(c *cli.Context) error {
			// create injector
			injector := do.DefaultInjector

			// listen to os interrupt signals and close the context
			ctx, cancel := signal.NotifyContext(c.Context, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
			defer cancel()

			ctx = extend.NewDelayedCancelContext(ctx, 5*time.Second)

			// inject the signal notify context
			do.ProvideValue(injector, ctx)

			// needed to use flags provided by the cmd.Run command
			c.Context = ctx
			do.OverrideValue(injector, c)

			provider.Connection(injector)
			do.Provide(injector, clinic.ProvideService)
			do.Provide(injector, clinic.NewClinicHandler)

			waitForTheEnd := &sync.WaitGroup{}

			waitForTheEnd.Add(1)
			go func() {
				defer waitForTheEnd.Done()

				router := appHttp.NewRouter()
				router.RegisterApplicationRoutes(
					do.MustInvoke[*clinic.ClinicHandler](injector),
				)

				httpServer := appHttp.NewServer(injector, router)
				go func() {
					<-ctx.Done()
					if err := httpServer.Shutdown(); err != nil {
						log.Printf("Failed to shutdown HTTP server: %v\n", err)
					}
				}()
				if err := httpServer.Start(); !errors.Is(err, http.ErrServerClosed) {
					log.Printf("Failed to start HTTP server: %v\n", err)
				}
				log.Println("Server has been stopped")
			}()

			// wait for context to be closed
			<-ctx.Done()

			waitForTheEnd.Wait()

			return injector.Shutdown()
		},
	}
}
