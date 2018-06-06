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

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/tasks"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//go:generate go run ./generate/generate.go
//go:generate go run ./generate/generate.go pluginParsers

var (
	version string
)

func main() {

	// Fix for #13, to provide a fallback version for plugin developers
	// In general it's recommended to use the official builds.
	if version == "" {
		version = "BUILT_FROM_SOURCE"
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Version = version
	app.Name = "kombustion"
	app.Usage = "Extend CloudFormation with plugins."
	app.Before = func(c *cli.Context) error {

		verbose := c.Bool("verbose")

		if verbose {
			log.SetLevel(log.InfoLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}

		// Init the spinner/printer
		printer.Init(verbose, "yellow", 14)

		// This is the initial loading message
		printer.Progress("Kombusting")
		return nil
	}

	app.Flags = tasks.GlobalFlags

	app.Commands = []cli.Command{
		// Manifest
		{
			Name:      "init",
			Usage:     "init manifest file",
			UsageText: "initialise a new manifest file in the current directory",
			Action:    tasks.InitaliseNewManifestTask,
			Flags:     tasks.InitManifestFlags,
		},
		// Plugin management
		{
			Name:      "add",
			Usage:     "add github.com/organisation/plugin",
			UsageText: "add github.com/organisation/plugin github.com/organisation/pluginTwo",
			Action:    tasks.AddPluginToManifest,
		},
		{
			Name:      "install",
			Usage:     "install all plugins in kombustion.yaml",
			UsageText: "install all plugins in kombustion.yaml",
			Action:    tasks.InstallPlugins,
		},
		// Cloudformation
		{
			Name:      "generate",
			Usage:     "parse a cloudformation template from ./config",
			UsageText: "kombustion cloudformation generate [command options] [stack]",
			Action:    tasks.Generate,
			Flags:     tasks.GenerateFlags,
		},
		{
			Name:      "upsert",
			Usage:     "upsert a cloudformation template or a yaml config",
			UsageText: "kombustion cloudformation upsert [command options] [stack]",
			Action:    tasks.Upsert,
			Flags:     tasks.UpsertFlags,
		},
		{
			Name:      "delete",
			Usage:     "delete a cloudformation stack",
			UsageText: "kombustion cloudformation delete [command options] [stackName]",
			Action:    tasks.Delete,
			Flags:     tasks.CloudFormationStackFlags,
		},
		{
			Name:      "events",
			Usage:     "print all events for a cloudformation stack ",
			UsageText: "kombustion cloudformation events [command options] [stackName]",
			Action:    tasks.PrintEvents,
			Flags:     tasks.CloudFormationStackFlags,
		},
	}

	app.Run(os.Args)
}
