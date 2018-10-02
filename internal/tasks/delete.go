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

	client := &cloudformation.Wrapper{}
	profile := c.GlobalString("profile")
	region := c.String("region")
	envName := c.String("environment")
	stackName := c.String("stack-name")
	manifestFile := c.GlobalString("manifest-file")

	taskDelete(
		client,
		objectStore,
		fileName,
		profile,
		stackName,
		region,
		envName,
		manifestFile,
	)
}

func taskDelete(
	client cloudformation.StackDeleter,
	objectStore core.ObjectStore,
	templatePath string,
	profileName string,
	stackName string,
	region string,
	envName string,
	manifestLocation string,
) {
	printer.Progress("Kombusting")

	manifestFile, err := manifest.GetManifestObject(objectStore, manifestLocation)
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

	acctID := client.Open(profileName, region)
	if env, ok := manifestFile.Environments[envName]; ok {
		if !env.IsAllowlistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, envName),
				"Use allowlisted account, or add account to environment accounts in kombustion.yaml",
				"",
			)
		}
	}

	tasks.DeleteStack(
		client,
		fullStackName,
		region,
	)
}
