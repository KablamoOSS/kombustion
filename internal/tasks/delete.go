package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

var DeleteFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to delete from",
		Value: "ap-southeast-2",
	},
}

func Delete(c *cli.Context) {
	tasks.DeleteStack(c.Args().Get(0), c.GlobalString("profile"), c.String("region"))
}
