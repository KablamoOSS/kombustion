package tasks

import "github.com/urfave/cli"

// GlobalFlags for all tasks
var GlobalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "output with high verbosity",
	},
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. ( -p Env=dev -p BucketName=test )",
	},
}

// CloudFormationStackFlags for tasks relating to CRUD of cloudformation stacks
var CloudFormationStackFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to deploy to",
	},
	cli.StringFlag{
		Name:  "stack-name",
		Usage: "stack name to deploy (defaults to ProjectName-Filename--Environment)",
	},
	cli.StringFlag{
		Name:  "environment, e",
		Usage: "environment config to use from ./kombustion.yaml",
		// Help:  "If you omit this, it will be derived based on the account the role is assumed from, provided that account is listed under an environment.",
	},
}
