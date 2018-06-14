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
		Usage: "cloudformation parameters `BucketName=test`",
	},
	cli.StringFlag{
		Name:  "profile",
		Usage: "use a profile from ~/.aws/credentials eg `MyProfile`",
	},
}

// CloudFormationStackFlags for tasks relating to CRUD of cloudformation stacks
var CloudFormationStackFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to deploy to eg `us-east-1`",
	},
	cli.StringFlag{
		Name:  "stack-name",
		Usage: "stack name to deploy [Default: ProjectName-Filename-Environment] eg `StackName-Environment`",
	},
	cli.StringFlag{
		Name:  "environment, e",
		Usage: "environment config to use from ./kombustion.yaml eg `production`",
	},
}
