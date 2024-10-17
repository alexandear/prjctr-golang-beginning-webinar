package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/samber/do"
	"github.com/urfave/cli/v2"
	"gocourse22/cmd"
	"gocourse22/cmd/flag"
	"log"
	"os"
	"runtime"
)

var (
	BuildVersion = "-"
	BuildBranch  = "-"
	BuildTag     = "-"
	BuildTime    = "-"
)

var info = fmt.Sprintf("build version: %s, branch: %s, tag: %s, build time: %s", BuildVersion, BuildBranch, BuildTag, BuildTime)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
	}

	runtime.SetBlockProfileRate(1)
}

// @title Clinics service
func main() {
	// create injector from default injector
	var injector = do.DefaultInjector

	// create the cli app
	app := &cli.App{
		Name:      "Clinics service",
		UsageText: "./app [global options] command [command options] [arguments...]\n",
		Version:   info,
		Flags: []cli.Flag{
			flag.EnvFileFlag(),
		},
		Authors: []*cli.Author{
			{
				Name:  "John Doe",
				Email: "john.doe@gmail.com",
			},
		},
		Commands: []*cli.Command{
			cmd.Run(),
			cmd.Migrate(),
			cmd.Worker(),
		},
		Before: func(c *cli.Context) error {
			// provide the cli context to the injector
			do.ProvideValue(injector, c)

			return nil
		},
	}

	// run the application
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
