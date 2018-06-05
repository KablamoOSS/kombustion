package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

// UpsertFlags for kombustion upsert
var UpsertFlags = []cli.Flag{
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. ( --param Env=dev --param BucketName=test )",
	},
	cli.BoolFlag{
		Name:  "no-base-outputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
	cli.BoolFlag{
		Name:  "iam, i",
		Usage: "gives the capability to perform upserts of IAM resources",
	},
}

func init() {
	UpsertFlags = append(CloudFormationStackFlags)
}

// Upsert a stack
func Upsert(c *cli.Context) {
	paramMap := getParamMap(c)

	generateParams := cloudformation.GenerateParams{
		Filename:           c.Args().Get(0),
		EnvFile:            c.String("env-file"),
		Env:                c.String("env"),
		DisableBaseOutputs: c.Bool("no-base-outputs"),
		ParamMap:           paramMap,
	})
	tasks.UpsertStack(generateParams, c.GlobalString("profile"), c.String("region"))
}
