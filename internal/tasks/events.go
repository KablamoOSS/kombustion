package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

// PrintEvents outputs the events of a stack
func PrintEvents(c *cli.Context) {
	cf := tasks.GetCloudformationClient(c.GlobalString("profile"), c.String("region"))
	stackName := c.Args().Get(0)
	tasks.PrintStackEvents(cf, stackName)
}
