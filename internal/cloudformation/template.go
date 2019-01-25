package cloudformation

import (
	"bytes"
	"fmt"
	"strings"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/pkg/parsers"
	"github.com/KablamoOSS/kombustion/types"
	"github.com/KablamoOSS/yaml"
)

func init() {
	registerYamlTagUnmarshalers()
}

// GenerateYamlTemplate - generate a cloudformation template
func GenerateYamlTemplate(params GenerateParams) (compiledTemplate YamlCloudformation, err error) {
	// load the config file
	var configData []byte

	// Setup out parser variables
	var templateParsers,
		coreParsers,
		pluginParsers map[string]types.ParserFunc

	templateParsers = make(map[string]types.ParserFunc)
	coreParsers = make(map[string]types.ParserFunc)
	pluginParsers = make(map[string]types.ParserFunc)

	// Load core AWS parsers for resources
	coreParsers = parsers.GetParsersResources()

	templateParsers = mergeParsers(templateParsers, coreParsers)

	// Load the parsers from Plugins
	pluginParsers = plugins.ExtractParsersFromPlugins(params.Plugins)
	templateParsers = mergeParsers(templateParsers, pluginParsers)

	if configData, err = params.ObjectStore.Get(params.Filename); err != nil {
		return compiledTemplate, err
	}

	//preprocess - template in the environment variables and custom params
	buf := new(bytes.Buffer)

	if err = executeTemplate(buf, configData, params.ParamMap); err != nil {
		printer.Error(
			fmt.Errorf("Failed to execute the template"),
			fmt.Sprintf(
				"File: %s",
				params.Filename,
			),
			"",
		)

		return compiledTemplate, err
	}

	// parse the config yaml
	data := buf.Bytes()
	var config YamlConfig

	config.Conditions = make(types.TemplateObject)
	config.Metadata = make(types.TemplateObject)
	config.Mappings = make(types.TemplateObject)
	config.Outputs = make(types.TemplateObject)
	config.Parameters = make(types.TemplateObject)
	config.Resources = make(map[string]types.CfResource)
	config.Transform = make(types.TemplateObject)

	if err = yaml.Unmarshal(data, &config); err != nil {
		printer.Error(
			fmt.Errorf("Failed to unmarshal the template"),
			fmt.Sprintf(
				"File: %s",
				params.Filename,
			),
			"",
		)
		return compiledTemplate, err
	}

	// Setup the initial types
	var conditions,
		metadata,
		mappings,
		outputs,
		parameters,
		resources,
		transform map[string]interface{}

	conditions = make(map[string]interface{})
	metadata = make(map[string]interface{})
	mappings = make(map[string]interface{})
	outputs = make(map[string]interface{})
	parameters = make(map[string]interface{})
	resources = make(map[string]interface{})
	transform = make(map[string]interface{})

	// Process the core and plugin parsers
	conditions,
		metadata,
		mappings,
		outputs,
		parameters,
		resources,
		transform = processParsers(config.Resources, templateParsers, params.GenerateDefaultOutputs)

	compiledTemplate = YamlCloudformation{
		AWSTemplateFormatVersion: config.AWSTemplateFormatVersion,
		Description:              config.Description,
		Metadata:                 mergeTemplates(config.Metadata, metadata),
		Parameters:               mergeTemplates(config.Parameters, parameters),
		Conditions:               mergeTemplates(config.Conditions, conditions),
		Transform:                mergeTemplates(config.Transform, transform),
		Mappings:                 mergeTemplates(config.Mappings, mappings),
		// As we've processed all the resources through our parsers, we don't want to merge in the
		// initial resources, as we will retain plugin definitions that don't map to
		// cloudformation template resources
		Resources: resources,
		Outputs:   mergeTemplates(config.Outputs, outputs),
	}

	return compiledTemplate, nil
}

// Process the parser funcs against the template's resources
// and return new template objects
func processParsers(
	templateResources types.ResourceMap,
	parserFuncs map[string]types.ParserFunc,
	generateDefaultOutputs bool,
) (
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
) {
	conditions = make(types.TemplateObject)
	metadata = make(types.TemplateObject)
	mappings = make(types.TemplateObject)
	outputs = make(types.TemplateObject)
	parameters = make(types.TemplateObject)
	resources = make(types.TemplateObject)
	transform = make(types.TemplateObject)

	// Loop through each Resource in the template, and parse it with a ParserFunc
	for templateResourceName, templateResource := range templateResources {
		var willMergeOutputs bool

		if strings.HasPrefix(templateResource.Type, "AWS::") {
			if generateDefaultOutputs {
				// if the flag is set merge the default generated outputs in
				willMergeOutputs = true
			} else {
				// If the flag is not set default behaviour is to ignore outputs
				willMergeOutputs = false
			}
		} else {
			// If this is a plugin merge the outputs
			willMergeOutputs = true
		}

		// If this is a custom resource, pass it through without touching it
		if templateResource.Type == "AWS::CloudFormation::CustomResource" ||
			strings.HasPrefix(templateResource.Type, "Custom::") {
			resources = mergeTemplatesWithError(
				templateResourceName,
				"non-plugin-resource",
				templateResource.Type,
				resources,
				types.TemplateObject{templateResourceName: templateResource},
			)
		} else if strings.HasPrefix(templateResource.Type, "AWS::") {
			resources = mergeTemplatesWithError(
				templateResourceName,
				"aws-resource",
				templateResource.Type,
				resources,
				types.TemplateObject{templateResourceName: templateResource},
			)
		} else { // This is a resource
			// Check if there is a parser for this resource
			parser, ok := parserFuncs[templateResource.Type]

			// If theres no parser dont adjust the out
			if !ok {
				resources = mergeTemplatesWithError(
					templateResourceName,
					"unknown-resource",
					templateResource.Type,
					resources,
					types.TemplateObject{templateResourceName: templateResource},
				)
				continue
			}

			// Marshal the resource into YAML to send to the parser function
			resourceData, err := yaml.Marshal(templateResource)

			if err != nil {
				return
			}

			parserSource,
				parserConditions,
				parserMetadata,
				parserMappings,
				parserOutputs,
				parserParameters,
				parserResources,
				parserTransform,
				parserErrors := parser(templateResourceName, string(resourceData))

			// If there were parser errors log them out
			if parserErrors != nil && parserErrors[0] != nil {
				for _, err := range parserErrors {
					parserError(
						err,
						templateResourceName,
						parserSource,
						templateResource.Type,
					)
				}
			}

			// Merge the results back together
			conditions = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				conditions,
				parserConditions,
			)

			metadata = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				metadata,
				parserMetadata,
			)

			mappings = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				mappings,
				parserMappings,
			)

			if willMergeOutputs {
				outputs = mergeTemplatesWithError(
					templateResourceName,
					parserSource,
					templateResource.Type,
					outputs,
					parserOutputs,
				)
			}

			parameters = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				parameters,
				parserParameters,
			)

			resources = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				resources,
				parserResources,
			)

			transform = mergeTemplatesWithError(
				templateResourceName,
				parserSource,
				templateResource.Type,
				transform,
				parserTransform,
			)
		}
	}
	return
}

// Merge Functions

func mergeParsers(maps ...map[string]types.ParserFunc) map[string]types.ParserFunc {
	result := make(map[string]types.ParserFunc)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func mergeTemplatesWithError(
	name,
	source,
	resourceType string,
	maps ...types.TemplateObject,
) types.TemplateObject {
	result := make(map[string]interface{})
	result = maps[0]
	for i, m := range maps {
		if i >= 1 { // map[0] is used as the starting point, so ignore it in this loop
			for k, v := range m {
				if _, exists := result[k]; !exists {
					result[k] = v
				} else {
					parserError(
						fmt.Errorf("Duplicate key for %s", k),
						name,
						source,
						resourceType,
					)
				}
			}
		}
	}
	return result
}

func mergeTemplates(maps ...types.TemplateObject) types.TemplateObject {
	result := make(types.TemplateObject)
	result = maps[0]
	for i, m := range maps {
		if i >= 1 { // map[0] is used as the starting point, so ignore it in this loop
			for k, v := range m {
				result[k] = v
			}
		}
	}
	return result
}

func mergeResources(
	configResources types.ResourceMap,
	baseResources types.TemplateObject,
) (
	combinedResource types.TemplateObject,
) {
	combinedResource = make(types.TemplateObject)

	// The yaml round-trip here is because there may be problems with the
	// msgpack round-trip from plugins. This is because we don't tag structs
	// with `msgpack:FieldName,omitempty` (doing so would also require plugin
	// writers to do the same). The yaml round-trip here ensures that
	// empty/null values are handled correctly.
	for k, v := range configResources {
		if obj, err := yaml.Marshal(v); err == nil {
			var tempCfResource types.CfResource
			if err = yaml.Unmarshal(obj, &tempCfResource); err == nil {
				combinedResource[k] = tempCfResource
			}
		}
	}

	for k, v := range baseResources {
		if obj, err := yaml.Marshal(v); err == nil {
			var tempCfResource types.CfResource
			if err = yaml.Unmarshal(obj, &tempCfResource); err == nil {
				combinedResource[k] = tempCfResource
			}
		}
	}

	return
}

// Print a error with a parser function
func parserError(err error, name, source, resourceType string) {
	printer.Error(
		err,
		fmt.Sprintf(
			"\n   ├─ Name:    %s\n   ├─ Source:  %s\n   └─ Type:    %s",
			name,
			source,
			resourceType,
		),
		"",
	)
}
