package flag

import (
	"github.com/urfave/cli/v2"
)

const (
	InstanceID  = "instance-id"
	InstanceEnv = "instance-env"
	VarDir      = "var-dir"
)

func InstanceEnvFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "Instance",
		Name:        InstanceEnv,
		Usage:       "Instance Environment",
		EnvVars:     []string{"ENV"},
		Value:       `dev`,
		DefaultText: `dev`,
	}
}

func InstanceIDFlag() cli.Flag {
	return &cli.IntFlag{
		Category:    "Instance",
		Name:        InstanceID,
		Usage:       "Instance ID",
		EnvVars:     []string{"INSTANCE_ID"},
		Value:       0,
		DefaultText: "0",
	}
}

func VarDirectoryFlag() cli.Flag {
	return &cli.StringFlag{
		Category:    "Instance",
		Name:        VarDir,
		Usage:       "Directory for variable content",
		EnvVars:     []string{"VAR_DIR"},
		DefaultText: "",
	}
}
