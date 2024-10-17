package flag

import (
	"github.com/urfave/cli/v2"
	"time"
)

const (
	HTTPServerAddress   = "http-server-address"
	HTTPReadTimeout     = "http-read-timeout"
	HTTPShutdownTimeout = "http-shutdown-timeout"
)

// HTTPServerAddressFlag is a flag to set http server address.
func HTTPServerAddressFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "http server",
		Name:        HTTPServerAddress,
		Usage:       "Set http server address",
		EnvVars:     []string{"HTTP_SERVER_ADDRESS"},
		DefaultText: ":8080",
		Value:       ":8080",
	}
}

// HTTPReadTimeoutFlag is a flag to set http read timeout.
func HTTPReadTimeoutFlag() cli.Flag {
	return &cli.DurationFlag{
		Category:    "http server",
		Name:        HTTPReadTimeout,
		Usage:       "Set http read timeout",
		EnvVars:     []string{"HTTP_READ_TIMEOUT"},
		DefaultText: "10s",
		Value:       10 * time.Second,
	}
}

// HTTPShutdownTimeoutFlag is a flag to set http shutdown timeout.
func HTTPShutdownTimeoutFlag() cli.Flag {
	return &cli.DurationFlag{
		Category:    "http server",
		Name:        HTTPShutdownTimeout,
		Usage:       "Set http shutdown timeout",
		EnvVars:     []string{"HTTP_SHUTDOWN_TIMEOUT"},
		DefaultText: "10s",
		Value:       10 * time.Second,
	}
}
