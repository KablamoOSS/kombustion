package manifest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestManifestLoad - Test unmarshalling the kombustion.yaml manifest
// into it's struct
//
// WARNING
// gofmt may try and convert the spaces in the yaml strings, into tabs
// If it doesn this, you will have a bad time. Tabs are banned characters
// in yaml.
func TestLoadManifestFromString(t *testing.T) {
	tests := []struct {
		name         string
		manifestYaml string
		output       Manifest
		throws       bool
	}{
		{
			name:         "Simple manifest",
			manifestYaml: `Name: TestManifest`,
			throws:       false,
			output: Manifest{
				Name:               "TestManifest",
				Region:             "",
				Plugins:            map[string]Plugin(nil),
				Architectures:      []string(nil),
				Environments:       map[string]Environment(nil),
				GenerateDefaultOutputs: false,
			},
		},
		{
			name: "Will throw error, due to tab character",
			manifestYaml: `
			name: TestManifest
`,
			throws: true,
		},
		{
			name: "Simple manifest GenerateDefaultOutputs",
			manifestYaml: `Name: TestManifest
GenerateDefaultOutputs: true`,
			output: Manifest{
				Name:               "TestManifest",
				Region:             "",
				Plugins:            map[string]Plugin(nil),
				Architectures:      []string(nil),
				Environments:       map[string]Environment(nil),
				GenerateDefaultOutputs: true,
			},
			throws: false,
		},
		{
			name:   "Manifest with github plugins",
			throws: false,
			manifestYaml: `Name: TestManifestWithPlugins
Region: us-east-1
Architectures: ["darwin/x64", "linux/386"]
Plugins:
  # Plugin 1 tests the latest version condition
  "github.com/KablamoOSS/kombustion-example-plugin-one@latest":
    Name: github.com/KablamoOSS/kombustion-example-plugin-one
    Version: latest

  # Plugin 2 tests the equals/exact version condition
  "github.com/KablamoOSS/kombustion-example-plugin-two@=2.0.1":
    Name: github.com/KablamoOSS/kombustion-example-plugin-two
    Version: "=2.0.1"

  # Plugin 3 tests the greater than version condition
  "github.com/KablamoOSS/kombustion-example-plugin-three@>3.x.x":
    Name: github.com/KablamoOSS/kombustion-example-plugin-three
    Version: ">3.x.x"
  # Plugin 4 tests the less than version condition
  "github.com/KablamoOSS/kombustion-example-plugin-four@<4.x.x":
    Name: github.com/KablamoOSS/kombustion-example-plugin-four
    Version: "<4.x.x"
`,
			output: Manifest{
				Name:          "TestManifestWithPlugins",
				Region:        "us-east-1",
				Architectures: []string{"darwin/x64", "linux/386"},
				Plugins: map[string]Plugin{
					"github.com/KablamoOSS/kombustion-example-plugin-one@latest": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-one",
						Version: "latest",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-two@=2.0.1": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-two",
						Version: "=2.0.1",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-three@>3.x.x": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-three",
						Version: ">3.x.x",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-four@<4.x.x": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-four",
						Version: "<4.x.x",
					},
				},
				Environments:       map[string]Environment(nil),
				GenerateDefaultOutputs: false,
			},
		},
		{
			name:   "Enviroment configuration",
			throws: false,
			manifestYaml: `Name: TestManifestWithEnvironment
Region: us-east-1
Environments:
  development:
    AccountIDs: [ "11111111111", "22222222222" ]
    Parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"

  staging:
    AccountIDs: [ "555555555"]
    Parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"

  production:
    AccountIDs: [ "55555555", "66666666"]
    Parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"
`,
			output: Manifest{
				Name:               "TestManifestWithEnvironment",
				Region:             "us-east-1",
				Plugins:            map[string]Plugin(nil),
				Architectures:      []string(nil),
				GenerateDefaultOutputs: false,
				Environments: map[string]Environment{
					"development": {
						AccountIDs: []string{"11111111111", "22222222222"},
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
					"staging": {
						AccountIDs: []string{"555555555"},
						Parameters: map[string]string{
							"parameterOneName":   "parameterOneValue",
							"parameterTwoName":   "8654238642489624862",
							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
						},
					},
					"production": {
						AccountIDs: []string{"55555555", "66666666"},
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
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput, err := loadManifestFromString([]byte(test.manifestYaml))
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
		}
	}
}

// FIXME: This test works locally, and fails in travis
// func TestFindAndLoadManifest(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		input  string
// 		output Manifest
// 		throws bool
// 	}{
// 		{
// 			name:  "Find manifest file in testdata",
// 			input: "testdata/works",
// 			output: Manifest{
// 				Name: "KombustionTest",
// 				Plugins: map[string]Plugin{
// 					"github.com/KablamoOSS/kombustion-example-plugin-one@latest": {
// 						Name:    "github.com/KablamoOSS/kombustion-example-plugin-one",
// 						Version: "latest",
// 					},
// 				},
// 				Architectures:      []string(nil),
// 				GenerateDefaultOutputs: false,
// 				Environments: map[string]Environment{
// 					"development": {
// 						AccountIDs: []string{"11111111111", "22222222222"},
// 						Parameters: map[string]string{
// 							"parameterOneName":   "parameterOneValue",
// 							"parameterTwoName":   "8654238642489624862",
// 							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
// 							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
// 						},
// 					},
// 					"staging": {
// 						AccountIDs: []string{"555555555"},
// 						Parameters: map[string]string{
// 							"parameterOneName":   "parameterOneValue",
// 							"parameterTwoName":   "8654238642489624862",
// 							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
// 							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
// 						},
// 					},
// 					"production": {
// 						AccountIDs: []string{"55555555", "66666666"},
// 						Parameters: map[string]string{
// 							"parameterOneName":   "parameterOneValue",
// 							"parameterTwoName":   "8654238642489624862",
// 							"parameterThreeName": "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s",
// 							"parameterFourName":  "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-",
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			name:   "Find manifest file in testdata",
// 			input:  "testdata/errors",
// 			throws: true,
// 		},
// 		{
// 			name:   "Find manifest file in testdata",
// 			input:  "testdata/empty",
// 			throws: true,
// 		},
// 	}

// 	for i, test := range tests {
// 		assert := assert.New(t)
// 		testOutput, err := findAndLoadManifest(test.input)
// 		if test.throws {
// 			assert.NotNil(err)
// 		} else {
// 			assert.Nil(err)
// 			assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.name))
// 		}
// 	}
// }
