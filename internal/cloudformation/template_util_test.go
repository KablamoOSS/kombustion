package cloudformation

import (
	"fmt"
	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

// This test confirms the merge works, but does not
// test the underlying functions that were merged
func TestMergeParsers(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	type Input struct {
		map1 map[string]types.ParserFunc
		map2 map[string]types.ParserFunc
	}

	tests := []struct {
		input  Input
		output []string
	}{
		{
			input: Input{
				map1: map[string]types.ParserFunc{
					"parser1": func(
						name string,
						data string,
					) (
						source string,
						conditions types.TemplateObject,
						metadata types.TemplateObject,
						mappings types.TemplateObject,
						outputs types.TemplateObject,
						parameters types.TemplateObject,
						resources types.TemplateObject,
						transform types.TemplateObject,
						errors []error,
					) {
						source = name
						return
					},
					"parser2": func(
						name string,
						data string,
					) (
						source string,
						conditions types.TemplateObject,
						metadata types.TemplateObject,
						mappings types.TemplateObject,
						outputs types.TemplateObject,
						parameters types.TemplateObject,
						resources types.TemplateObject,
						transform types.TemplateObject,
						errors []error,
					) {
						source = name
						return
					},
				},
				map2: map[string]types.ParserFunc{
					"parser3": func(
						name string,
						data string,
					) (
						source string,
						conditions types.TemplateObject,
						metadata types.TemplateObject,
						mappings types.TemplateObject,
						outputs types.TemplateObject,
						parameters types.TemplateObject,
						resources types.TemplateObject,
						transform types.TemplateObject,
						errors []error,
					) {
						source = name
						return
					},
					"parser4": func(
						name string,
						data string,
					) (
						source string,
						conditions types.TemplateObject,
						metadata types.TemplateObject,
						mappings types.TemplateObject,
						outputs types.TemplateObject,
						parameters types.TemplateObject,
						resources types.TemplateObject,
						transform types.TemplateObject,
						errors []error,
					) {
						source = name
						return
					},
				},
			},
			output: []string{
				"parser1",
				"parser2",
				"parser3",
				"parser4",
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		output := mergeParsers(
			test.input.map1,
			test.input.map2,
		)

		matches := 0
		for key := range output {
			for _, outputKey := range test.output {
				if key == outputKey {
					matches = matches + 1
				}
			}
		}
		assert.Equal(
			len(test.output),
			matches,
			fmt.Sprintf("Test: %d", i),
		)
	}
}

func TestMergeFinalTemplatesWithErrors(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	type Input struct {
		name,
		source,
		resourceType string
		map1 types.TemplateObject
		map2 types.TemplateObject
	}

	tests := []struct {
		input  Input
		output types.TemplateObject
	}{
		{
			input: Input{
				name:         "test-1",
				source:       "test-case",
				resourceType: "test-case",
				map1: types.TemplateObject{
					"Test": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue",
						},
						DependsOn: []interface{}{
							"OtherResource",
						},
					},
				},
				map2: types.TemplateObject{
					"Test2": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup2",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue2",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue2",
						},
						DependsOn: []interface{}{
							"OtherResource2",
						},
					},
				},
			},
			output: types.TemplateObject{
				"Test": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue",
					},
					DependsOn: []interface{}{
						"OtherResource",
					},
				},
				"Test2": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup2",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue2",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue2",
					},
					DependsOn: []interface{}{
						"OtherResource2",
					},
				},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		output := mergeTemplatesWithError(
			test.input.name,
			test.input.source,
			test.input.resourceType,
			test.input.map1,
			test.input.map2,
		)

		assert.Equal(
			test.output,
			output,
			fmt.Sprintf("Test: %d", i),
		)
	}
}

func TestMergeFinalTemplates(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	type Input struct {
		map1 types.TemplateObject
		map2 types.TemplateObject
	}

	tests := []struct {
		input  Input
		output types.TemplateObject
	}{
		{
			input: Input{
				map1: types.TemplateObject{
					"Test": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue",
						},
						DependsOn: []interface{}{
							"OtherResource",
						},
					},
				},
				map2: types.TemplateObject{
					"Test2": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup2",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue2",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue2",
						},
						DependsOn: []interface{}{
							"OtherResource2",
						},
					},
				},
			},
			output: types.TemplateObject{
				"Test": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue",
					},
					DependsOn: []interface{}{
						"OtherResource",
					},
				},
				"Test2": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup2",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue2",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue2",
					},
					DependsOn: []interface{}{
						"OtherResource2",
					},
				},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		output := mergeTemplates(
			test.input.map1,
			test.input.map2,
		)

		assert.Equal(
			test.output,
			output,
			fmt.Sprintf("Test: %d", i),
		)
	}
}

func TestMergeResources(t *testing.T) {
	// Prevent the printer from exiting
	printer.Test()

	type Input struct {
		configResources types.ResourceMap
		resources       types.TemplateObject
	}

	tests := []struct {
		input  Input
		output types.TemplateObject
	}{
		{
			input: Input{
				configResources: types.ResourceMap{
					"Test": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue",
						},
						DependsOn: []interface{}{
							"OtherResource",
						},
					},
				},
				resources: types.TemplateObject{
					"Test2": types.CfResource{
						Type: "AWS::Logs::LogGroup",
						Properties: map[string]interface{}{
							"LogGroupName": "TestLogGroup2",
						},
						Condition: map[string]interface{}{
							"ConditionName": "ConditionValue2",
						},
						Metadata: map[string]interface{}{
							"MetadataName": "MetadataValue2",
						},
						DependsOn: []interface{}{
							"OtherResource2",
						},
					},
				},
			},
			output: types.TemplateObject{
				"Test": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue",
					},
					DependsOn: []interface{}{
						"OtherResource",
					},
				},
				"Test2": types.CfResource{
					Type: "AWS::Logs::LogGroup",
					Properties: map[string]interface{}{
						"LogGroupName": "TestLogGroup2",
					},
					Condition: map[string]interface{}{
						"ConditionName": "ConditionValue2",
					},
					Metadata: map[string]interface{}{
						"MetadataName": "MetadataValue2",
					},
					DependsOn: []interface{}{
						"OtherResource2",
					},
				},
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		output := mergeResources(
			test.input.configResources,
			test.input.resources,
		)

		assert.Equal(
			test.output,
			output,
			fmt.Sprintf("Test: %d", i),
		)
	}
}
