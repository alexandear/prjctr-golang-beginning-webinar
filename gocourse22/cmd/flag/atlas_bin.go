package flag

import "github.com/urfave/cli/v2"

const (
	AtlasBin = "atlas-bin"
)

func AtlasBinFlag() cli.Flag {
	return &cli.StringFlag{
		Name:        AtlasBin,
		Usage:       "Atlas binary",
		EnvVars:     []string{"ATLAS_BIN"},
		DefaultText: "atlas",
		Value:       "atlas",
	}
}
