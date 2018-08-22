package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/urfave/cli"
)

// DeleteFlags for use with the delete taks
var DeleteFlags = []cli.Flag{}

func init() {
	DeleteFlags = append(CloudFormationStackFlags, DeleteFlags...)
}

// Delete a given stack
func Delete(c *cli.Context) {
	printer.Progress("Kombusting")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't upsert file, no source template provided"),
			"Add the path to the source template file you want to generate like: `kombustion upsert template.yaml`.",
			"https://www.kombustion.io/api/manifest/",
		)
	}

	manifestFile := manifest.FindAndLoadManifest()

	environment := c.String("environment")

	stackName := cloudformation.GetStackName(manifestFile, fileName, environment, c.String("stack-name"))

	region := c.String("region")
	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}

	acctID, cf := tasks.GetCloudformationClient(c.GlobalString("profile"), region)
	if env, ok := manifestFile.Environments[environment]; ok {
		if !env.IsWhitelistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, environment),
				"Use whitelisted account, or add account to environment accounts in kombustion.yaml",
				"",
			)
		}
	}

	tasks.DeleteStack(
		cf,
		stackName,
		region,
	)
}
