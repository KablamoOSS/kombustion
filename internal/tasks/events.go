package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/urfave/cli"
)

// PrintEvents outputs the events of a stack
func PrintEvents(c *cli.Context) {
	objectStore := core.NewFilesystemStore(".")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't find stack, no source template provided"),
			"Add the path to the source template file, or provide use --stack-name",
			"https://www.kombustion.io/api/cli/#stacks",
		)
	}

	eventer := &cloudformation.Wrapper{}

	printEvents(
		objectStore,
		eventer,
		fileName,
		c.String("stack-name"),
		c.GlobalString("environment"),
		c.GlobalString("region"),
		c.String("profile"),
	)
}

func printEvents(
	objectStore core.ObjectStore,
	eventer cloudformation.StackEventer,
	templatePath string,
	stackName string,
	envName string,
	profile string,
	region string,
) {
	printer.Progress("Kombusting")

	manifestFile, err := manifest.GetManifestObject(objectStore)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}
	acctID := eventer.Open(profile, region)

	if env, ok := manifestFile.Environments[envName]; ok {
		if !env.IsAllowlistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, envName),
				"Use allowlisted account, or add account to environment accounts in kombustion.yaml",
				"",
			)
		}
	}

	stackName = cloudformation.GetStackName(
		manifestFile,
		templatePath,
		envName,
		stackName,
	)

	tasks.PrintStackEvents(eventer, stackName)
}
