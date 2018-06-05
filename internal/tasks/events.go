package tasks

import (
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/urfave/cli"
)

// PrintStackEvents outputs the events of a stack
func PrintEvents(c *cli.Context) {
	cf := getCF(c.GlobalString("profile"), c.String("region"))
	stackName := c.Args().Get(0)
	tasks.PrintStackEvents(cf, stackName)
}
