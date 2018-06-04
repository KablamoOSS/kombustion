// AWS API and CloudFormation parser.
//
// Installation
//
// Get the package
//
//     go get github.com/KablamoOSS/kombustion
//
// Build as docker image
//
//     docker build -t kombustion .
//
// Update Base Parsers
//
//     go run ./generate/generate.go
//
// Usage
//
// Generate a cloudformation template from: ./configs/test.yaml:
//
//     kombustion cf generate --format=yaml test
//
// Upsert a cloudformation template from: ./compiled/test.yaml:
//
//     kombustion cf upsert test
//
// Delete a cloudformation stack (stackName: test)
//
//     kombustion cf delete test
//
// Print all the events for a stack (stackName: test)
//
//     kombustion cf events test
//
// Using Roles and MFA
//
//     TOKEN=000000 \
//     MFA_SERIAL=arn:aws:iam::123456:mfa/stackCreator \
//     ASSUMED_ROLE=arn:aws:iam::123456:role/stackCreatorRole \
//     kombustion cf upsert test
//
// Custom Plugins
//
// Kombustion utilizes package plugin (https://godoc.org/plugin).
// By default kombustion will look for plugins in the ./plugins directory.
// You can also Specify custom plugins directory:
//
//     PLUGINS=/plugins kombustion cf generate test
//
package main

import (
	"os"

	"github.com/KablamoOSS/kombustion/manifest"
	"github.com/KablamoOSS/kombustion/plugins"
	"github.com/KablamoOSS/kombustion/tasks"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//go:generate go run ./generate/generate.go
//go:generate go run ./generate/generate.go pluginParsers

var (
	version string
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = version
	app.Name = "kombustion"
	app.Usage = "Extend CloudFormation with plugins."
	app.Before = func(c *cli.Context) error {
		log.SetLevel(log.WarnLevel)
		if c.Bool("verbose") {
			log.SetLevel(log.InfoLevel)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "output with high verbosity",
		},
	}

	app.Commands = []cli.Command{
		// Manifest
		{
			Name:      "init",
			Usage:     "init manifest file",
			UsageText: "initialise a new manifest file in the current directory",
			Action:    manifest.InitaliseNewManifestTask,
			// Flags not yet programmed
			// Flags:     manifest.InitManifestFlags,
		},
		// Plugin management
		{
			Name:      "add",
			Usage:     "add github.com/organisation/plugin",
			UsageText: "add github.com/organisation/plugin github.com/organisation/pluginTwo",
			Action:    plugins.AddPluginToManifest,
		},
		{
			Name:      "install",
			Usage:     "install all plugins in kombustion.yaml",
			UsageText: "install all plugins in kombustion.yaml",
			Action:    plugins.InstallPlugins,
		},
		// Cloudformation
		{
			Name:    "cloudformation",
			Aliases: []string{"cf"},
			Usage:   "tasks for building and deploying cloudformation templates",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "profile",
					Usage: "aws credentials profile to use",
				},
			},
			Subcommands: []cli.Command{
				{
					Name:      "generate",
					Usage:     "parse a cloudformation template from ./config",
					UsageText: "kombustion cloudformation generate [command options] [stack]",
					Action:    tasks.Generate,
					Flags:     tasks.Generate_Flags,
				},
				{
					Name:      "upsert",
					Usage:     "upsert a cloudformation template or a yaml config",
					UsageText: "kombustion cloudformation upsert [command options] [stack]",
					Action:    tasks.Upsert,
					Flags:     tasks.Upsert_Flags,
				},
				{
					Name:      "delete",
					Usage:     "delete a cloudformation stack",
					UsageText: "kombustion cloudformation delete [command options] [stackName]",
					Action:    tasks.Delete,
					Flags:     tasks.Delete_Flags,
				},
				{
					Name:      "events",
					Usage:     "print all events for a cloudformation stack ",
					UsageText: "kombustion cloudformation events [command options] [stackName]",
					Action:    tasks.PrintEvents,
					Flags:     tasks.PrintEvents_Flags,
				},
				{
					Name:      "plugins",
					Usage:     "get or list plugins (see cf plugins help)",
					UsageText: "kombustion cloudformation plugins [command options]",
					Subcommands: []cli.Command{
						{
							Name:      "get",
							Usage:     "install a plugin from the plugin repository",
							UsageText: "kombustion cloudformation plugins get [command options] pluginname",
							Action:    tasks.GetPlugin,
							Flags:     tasks.GetPlugin_Flags,
						},
						{
							Name:      "list",
							Usage:     "list all loaded plugins",
							UsageText: "kombustion cloudformation plugins list [command options]",
							Action:    tasks.PrintPlugins,
							Flags:     tasks.PrintPlugins_Flags,
						},
						{
							Name:      "delete",
							Usage:     "deletes the given plugin",
							UsageText: "kombustion cloudformation plugins delete [command options] pluginname",
							Action:    tasks.DeletePlugin,
							Flags:     tasks.DeletePlugin_Flags,
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
