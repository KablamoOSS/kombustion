package manifest

import (
	"fmt"
	"testing"
	// "github.com/google/go-cmp/cmp"
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
		throws       error
	}{
		{
			name:         "Simple manifest",
			manifestYaml: `name: TestManifest`,
			output: Manifest{
				Name:               "TestManifest",
				Plugins:            map[string]ManifestPlugin(nil),
				Architectures:      []string(nil),
				Environments:       map[string]ManifestEnvironment(nil),
				HideDefaultExports: false,
			},
		},
		{
			name: "Will throw error, due to tab character",
			manifestYaml: `
			name: TestManifest
`,
			throws: fmt.Errorf("line 2: found character that cannot start any token"),
		},
		{
			name: "Simple manifest HideDefaultExports",
			manifestYaml: `name: TestManifest
hideDefaultExports: true`,
			output: Manifest{
				Name:               "TestManifest",
				Plugins:            map[string]ManifestPlugin(nil),
				Architectures:      []string(nil),
				Environments:       map[string]ManifestEnvironment(nil),
				HideDefaultExports: true,
			},
		},
		{
			name: "Manifest with github plugins",
			manifestYaml: `name: TestManifestWithPlugins
architectures: ["darwin/x64", "linux/386"]
plugins:
  # Plugin 1 tests the latest version condition
  - name: github.com/KablamoOSS/kombustion-example-plugin-one
    version: latest

  # Plugin 2 tests the equals/exact version condition
  - name: github.com/KablamoOSS/kombustion-example-plugin-two
    version: "=2.0.1"

  # Plugin 3 tests the greater than version condition
  - name: github.com/KablamoOSS/kombustion-example-plugin-three
    version: ">3.x.x"

  # Plugin 4 tests the less than version condition
  - name: github.com/KablamoOSS/kombustion-example-plugin-four
    version: "<4.x.x"
`,
			output: Manifest{
				Name:          "TestManifestWithPlugins",
				Architectures: []string{"darwin/x64", "linux/386"},
				Plugins: map[string]ManifestPlugin{
					"github.com/KablamoOSS/kombustion-example-plugin-one": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-one",
						Version: "latest",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-two": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-two",
						Version: "=2.0.1",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-three": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-three",
						Version: ">3.x.x",
					},
					"github.com/KablamoOSS/kombustion-example-plugin-four": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-four",
						Version: "<4.x.x",
					},
				},
				Environments:       map[string]ManifestEnvironment(nil),
				HideDefaultExports: false,
			},
		},
		{
			name: "Enviroment configuration",
			manifestYaml: `name: TestManifestWithEnvironment
environments:
  development:
    accountIDs: [ "11111111111", "22222222222" ]
    parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"

  staging:
    accountIDs: [ "555555555"]
    parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"

  production:
    accountIDs: [ "55555555", "66666666"]
    parameters:
      parameterOneName: "parameterOneValue"
      parameterTwoName: "8654238642489624862"
      parameterThreeName: "3so87tg4y98n7y34ts3t4sh  st34y79p4y3t7 8s"
      parameterFourName: "hhh:://asdfasdf.sadfasdf:3452345@f][a;v-][0[-"
`,
			output: Manifest{
				Name:               "TestManifestWithEnvironment",
				Plugins:            map[string]ManifestPlugin(nil),
				Architectures:      []string(nil),
				HideDefaultExports: false,
				Environments: map[string]ManifestEnvironment{
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

	for _, test := range tests {
		_, err := loadManifestFromString([]byte(test.manifestYaml))
		if err != nil {
			if test.throws != nil {
				// currently not testing the error that is thrown, just that one is
			} else {
				t.Error(err)
			}
		}
		// if cmp.Equal(testOutput, test.output) == false {
		// 	if diff := cmp.Diff(testOutput, test.output); diff != "" {
		// 		t.Errorf("Test #%d [%s] output: (-got +want)\n%s", i, test.name, diff)
		// 	}
		// }
	}
}

func TestFindAndLoadManifest(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output Manifest
		throws error
	}{
		{
			name:  "Find manifest file in test-data",
			input: "test-data/works",
			output: Manifest{
				Name: "KombustionTest",
				Plugins: map[string]ManifestPlugin{
					"github.com/KablamoOSS/kombustion-example-plugin-one": {
						Name:    "github.com/KablamoOSS/kombustion-example-plugin-one",
						Version: "latest",
					},
				},
				Architectures:      []string(nil),
				HideDefaultExports: false,
				Environments: map[string]ManifestEnvironment{
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
		{
			name:   "Find manifest file in test-data",
			input:  "test-data/errors",
			throws: fmt.Errorf("there are both kombustion.yaml && kombustion.yml files, please remove one"),
		},
		{
			name:   "Find manifest file in test-data",
			input:  "test-data/empty",
			throws: fmt.Errorf("no kombustion.yaml manifest file found"),
		},
	}

	for _, test := range tests {
		_, err := findAndLoadManifest(test.input)
		if err != nil {
			if test.throws != nil {
				// currently not testing the error that is thrown, just that one is
			} else {
				t.Error(err)
			}
		}
		// if cmp.Equal(testOutput, test.output) == false {
		// 	if diff := cmp.Diff(testOutput, test.output); diff != "" {
		// 		t.Errorf("Test #%d [%s] output: (-got +want)\n%s", i, test.name, diff)
		// 	}
		// }
	}
}
