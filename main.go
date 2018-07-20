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
	"fmt"
	"os"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/tasks"
	log "github.com/sirupsen/logrus"
	"github.com/ttacon/chalk"
	"github.com/urfave/cli"
)

//go:generate go run ./generate/main.go ./generate/generate.go ./generate/cloudformationSpec.go ./generate/templates.go ./generate/types.go ./generate/util.go

var (
	version string
)

func main() {

	// Fix for #13, to provide a fallback version for plugin developers
	// In general it's recommended to use the official builds.
	if version == "" {
		version = "BUILT_FROM_SOURCE"

		devModeFlags := []cli.Flag{
			cli.StringFlag{
				Name:  "load-plugin",
				Usage: "load arbitrary plugin `path/to/plugin.so`",
			},
		}

		tasks.GlobalFlags = append(tasks.GlobalFlags, devModeFlags...)

	}

	kombustionLogo := chalk.Dim.TextStyle(`
   __              __            __  _
  / /_____  __ _  / /  __ _____ / /_(_)__  ___
 /  '_/ _ \/  ' \/ _ \/ // (_-</ __/ / _ \/ _ \
/_/\_\\___/_/_/_/_.__/\_,_/___/\__/_/\___/_//_/
kombustion.io
_______________________________________________________________________
`)

	cli.AppHelpTemplate = fmt.Sprintf(`%s
%s
ISSUES:
    If you have an issue with kombustion, check both the kombustion.io documentation [0], and
    the CloudFormation documentation [1] to help you resolve it.

    If the issue still persists please check out the issue queue [3] to see if it's
    already been reported and/or has a fix. If it hasn't you can create a new one [3].
%s`,
		kombustionLogo,
		cli.AppHelpTemplate,
		chalk.Dim.TextStyle(`
    [0] https://kombustion.io
    [1] https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-reference.html
    [3] https://github.com/KablamoOSS/kombustion/issues
`),
	)

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
		printer.Init(verbose, "yellow", 14, os.Stdout)
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
		// CloudFormation
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
		// Core - Self update
		{
			Name:   "update",
			Usage:  "update kombustion",
			Action: tasks.Update,
			Flags:  tasks.UpdateFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		if err.Error() == "flag provided but not defined: -load-plugin" {
			printer.Fatal(
				err,
				"--load-plugin is only available when kombustion is built from source. See the link below for more information.",
				"https://www.kombustion.io/guides/plugins/",
			)
		}
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}
}
