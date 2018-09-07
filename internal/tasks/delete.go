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

// DeleteFlags for use with the delete taks
var DeleteFlags = []cli.Flag{}

func init() {
	DeleteFlags = append(CloudFormationStackFlags, DeleteFlags...)
}

// Delete a given stack
func Delete(c *cli.Context) {
	objectStore := core.NewFilesystemStore(".")

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't upsert file, no source template provided"),
			"Add the path to the source template file you want to generate like: `kombustion upsert template.yaml`.",
			"https://www.kombustion.io/api/manifest/",
		)
	}

	profile := c.GlobalString("profile")
	region := c.String("region")
	envName := c.String("env")
	stackName := c.String("stackName")

	taskDelete(
		objectStore,
		fileName,
		profile,
		stackName,
		region,
		envName,
	)
}

func taskDelete(
	objectStore core.ObjectStore,
	templatePath string,
	profileName string,
	stackName string,
	region string,
	envName string,
) {
	printer.Progress("Kombusting")

	manifestFile, err := manifest.GetManifestObject(objectStore)
	if err != nil {
		printer.Fatal(err, config.ErrorHelpInfo, "")
	}

	fullStackName := cloudformation.GetStackName(manifestFile, templatePath, envName, stackName)

	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}

	acctID, cf := tasks.GetCloudformationClient(profileName, region)
	if env, ok := manifestFile.Environments[envName]; ok {
		if !env.IsWhitelistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, envName),
				"Use whitelisted account, or add account to environment accounts in kombustion.yaml",
				"",
			)
		}
	}

	tasks.DeleteStack(
		cf,
		fullStackName,
		region,
	)
}
