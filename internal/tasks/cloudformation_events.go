package tasks

import (
	"github.com/urfave/cli"
)

var PrintEvents_Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to deploy to",
		Value: "ap-southeast-2",
	},
}

func PrintEvents(c *cli.Context) {
	cf := getCF(c.GlobalString("profile"), c.String("region"))
	stackName := c.Args().Get(0)
	printStackEvents(cf, stackName)
}
