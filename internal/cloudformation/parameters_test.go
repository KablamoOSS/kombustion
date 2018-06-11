package cloudformation

import (
	"flag"
	"fmt"
	"testing"

	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/types"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

type mockCloudFormationClient struct {
	cloudformationiface.CloudFormationAPI
}

func TestResolveEnvironmentParameters(t *testing.T) {
	tests := []struct {
		name        string
		environment string
		manifest    manifest.Manifest
		output      map[string]string
	}{
		{
			name:        "Returns map of env vars",
			environment: "development",
			manifest: manifest.Manifest{
				Name:               "TestManifestWithEnvironment",
				Plugins:            nil,
				Architectures:      []string(nil),
				HideDefaultExports: false,
				Environments: map[string]manifest.Environment{
					"development": {
						AccountIDs: nil,
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
				},
			},
			output: map[string]string{
				"parameterOneName":   "parameterOneValue",
				"parameterTwoName":   "8654238642489624862",
				"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
				"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
			},
		},
		{
			name:        "Returns emtpy map",
			environment: "production",
			manifest: manifest.Manifest{
				Name:               "TestManifestWithEnvironment",
				Plugins:            nil,
				Architectures:      []string(nil),
				HideDefaultExports: false,
				Environments: map[string]manifest.Environment{
					"development": {
						AccountIDs: nil,
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
				},
			},
			output: nil,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := resolveEnvironmentParameters(&test.manifest, test.environment)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
	}
}

func TestResolveParameters(t *testing.T) {
	type input struct {
		ctx          *cli.Context
		cfYaml       YamlCloudformation
		cfClient     *awsCF.CloudFormation
		manifestFile *manifest.Manifest
	}

	tests := []struct {
		name   string
		input  input
		output []*awsCF.Parameter
	}{
		{
			name: "Dev",
			input: input{
				ctx: func() *cli.Context {
					set := flag.NewFlagSet("test", 0)
					set.String("environment", "development", "development")
					context := cli.NewContext(nil, set, nil)
					return context
				}(),
				cfYaml: YamlCloudformation{
					AWSTemplateFormatVersion: "version",
					Description:              "Test Template",
					Parameters: types.TemplateObject{
						"parameterOneName":   "parameterOneValue",
						"parameterTwoName":   "8654238642489624862",
						"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
						"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
					},
					Mappings:   types.TemplateObject{},
					Conditions: types.TemplateObject{},
					Transform:  types.TemplateObject{},
					Resources:  types.TemplateObject{},
					Outputs:    types.TemplateObject{},
				},
				manifestFile: &manifest.Manifest{
					Name:               "TestManifestWithEnvironment",
					Plugins:            nil,
					Architectures:      []string(nil),
					HideDefaultExports: false,
					Environments: map[string]manifest.Environment{
						"development": {
							AccountIDs: nil,
							Parameters: map[string]string{
								"parameterOneName":   "parameterOneValue",
								"parameterTwoName":   "8654238642489624862",
								"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
								"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
							},
						},
					},
				},
			},
			output: []*awsCF.Parameter{},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := ResolveParameters(
			test.input.ctx,
			test.input.cfYaml,
			test.input.manifestFile,
		)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
	}
}

// func ResolveParameters(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input struct {
// 			ctx          *cli.Context
// 			cfYaml       YamlCloudformation
// 			cfClient     *awsCF.CloudFormation
// 			manifestFile *manifest.Manifest
// 		}
// 		output []*awsCF.Parameter
// 		throws bool
// 	}{
// 		{
// 			name:        "Dev",
// 			environment: "development",
// 			output:      map[string]string{},
// 			throws:      false,
// 		},
// 	}

// 	for i, test := range tests {
// 		assert := assert.New(t)
// 		testOutput, err := ResolveParameters(
// 			test.input.ctx,
// 			test.input.cfYaml,
// 			test.input.cfClient,
// 			test.input.manifestFile,
// 		)
// 		if test.throws {
// 			assert.NotNil(err)
// 		} else {
// 			assert.Nil(err)
// 			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
// 		}
// 	}
// }
