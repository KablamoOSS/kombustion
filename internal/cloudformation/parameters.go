package cloudformation

import (
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

// ResolveParameters for the template
func ResolveParameters(
	environment string,
	cliParams map[string]string,
	cfYaml YamlCloudformation,
	manifestFile *manifest.Manifest,
) []*awsCF.Parameter {
	results := []*awsCF.Parameter{}

	env := resolveEnvironmentParameters(manifestFile, environment)

	if env == nil {
		env = make(map[string]string)
	}

	// override envFile values with optional --param values
	for key, value := range cliParams {
		env[key] = value
	}

	// Filter all available Parameters to only those present
	// in the template
	for paramKey := range cfYaml.Parameters {
		for key, value := range env {
			if paramKey == key {
				// Filter to params in the stack
				results = append(results, &awsCF.Parameter{
					ParameterKey:   aws.String(key),
					ParameterValue: aws.String(value),
				})
			}
		}
	}
	return results
}

// ResolveParametersS3 for an S3 based template
func ResolveParametersS3(
	c *cli.Context,
	manifestFile *manifest.Manifest,
) []*awsCF.Parameter {

	results := []*awsCF.Parameter{}

	params := make(map[string]string)

	env := resolveEnvironmentParameters(manifestFile, c.String("environment"))

	// convert to aws Parameter list
	// TODO: We probably need to download the template to determine what params
	// it needs, and filter the available params only to those

	for key, value := range params {
		// Filter to params in the stack
		results = append(results, &awsCF.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		})
	}

	for key, value := range env {
		results = append(results, &awsCF.Parameter{
			ParameterKey:   aws.String(key),
			ParameterValue: aws.String(value),
		})
	}

	return results
}

func resolveEnvironmentParameters(manifestFile *manifest.Manifest, environment string) (parameters map[string]string) {
	if manifestFile.Environments[environment].Parameters != nil {
		envParams := manifestFile.Environments[environment].Parameters
		if envParams != nil {
			parameters = envParams
		}
	}
	return
}
