package flag

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	EnvFile = "env"
)

func EnvFileFlag() cli.Flag {
	return &cli.StringFlag{
		Name:        EnvFile,
		Aliases:     []string{"e"},
		Usage:       "path to env file",
		DefaultText: ".env",
		Value:       ".env",

		// override environment variables if the option is present
		Action: func(context *cli.Context, envFile string) error {
			return godotenv.Load(envFile)
		},
	}
}
