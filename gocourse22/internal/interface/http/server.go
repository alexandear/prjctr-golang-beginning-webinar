package http

import (
	"context"
	"net/http"
	"time"

	"github.com/samber/do"
	"github.com/urfave/cli/v2"

	"prjctr.com/gocourse22/cmd/flag"
)

// HTTP is the http server
type HTTP struct {
	srv             *http.Server
	shutdownTimeout time.Duration
}

// NewHTTP creates a new server
func NewHTTP(srv *http.Server, st time.Duration) *HTTP {
	return &HTTP{srv: srv, shutdownTimeout: st}
}

// Start starts the server
func (s HTTP) Start() error {
	return s.srv.ListenAndServe()
}

// Stop gracefully stops the server
// https://golang.org/pkg/net/http/#Server.Shutdown
func (s HTTP) Stop(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s HTTP) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.Stop(ctx)
}

// NewServer provides a new http server
func NewServer(injector *do.Injector, router *Router) *HTTP {
	c := do.MustInvoke[*cli.Context](injector)

	serv := &http.Server{
		Addr:              c.String(flag.HTTPServerAddress),
		Handler:           router.Handler(),
		ReadTimeout:       c.Duration(flag.HTTPReadTimeout),
		ReadHeaderTimeout: c.Duration(flag.HTTPReadTimeout),
	}

	return NewHTTP(serv, c.Duration(flag.HTTPShutdownTimeout))
}
