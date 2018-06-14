package tasks

import (
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

// DeleteFlags for use with the delete taks
var DeleteFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to delete from",
		Value: "ap-southeast-2",
	},
}

// Delete a given stack
func Delete(c *cli.Context) {
	printer.Step("Deleting stack")
	printer.Progress("Kombusting")

	tasks.DeleteStack(c.Args().Get(0), c.GlobalString("profile"), c.String("region"))
}
