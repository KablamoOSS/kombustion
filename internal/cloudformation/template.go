package cloudformation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	printer "github.com/KablamoOSS/go-cli-printer"
	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/pkg/parsers"
	"github.com/KablamoOSS/kombustion/types"

	"github.com/KablamoOSS/yaml"
)

func init() {
	registerYamlTagUnmarshalers()
}

// GenerateYamlStack - generate a stack
func GenerateYamlStack(params GenerateParams) (compiledTemplate YamlCloudformation, err error) {
	// load the config file
	var configData []byte

	// Setup out parser variables
	var templateParsers,
		coreParsers,
		pluginParsers,
		coreOutputParsers map[string]types.ParserFunc

	templateParsers = make(map[string]types.ParserFunc)
	coreParsers = make(map[string]types.ParserFunc)
	pluginParsers = make(map[string]types.ParserFunc)
	coreOutputParsers = make(map[string]types.ParserFunc)

	// Load core AWS parsers for resources
	coreParsers = parsers.GetParsersResources()

	templateParsers = mergeParsers(templateParsers, coreParsers)

	// If we're generating outputs, load the output parsers
	if params.GenerateDefaultOutputs {
		coreOutputParsers = parsers.GetParsersOutputs()
		templateParsers = mergeParsers(templateParsers, coreOutputParsers)
	}

	// Load the parsers from Plugins
	pluginParsers = plugins.ExtractParsersFromPlugins(params.Plugins)
	templateParsers = mergeParsers(templateParsers, pluginParsers)

	if configData, err = ioutil.ReadFile(params.Filename); err != nil {
		return compiledTemplate, err
	}

	//preprocess - template in the environment variables and custom params
	buf := new(bytes.Buffer)

	if err = executeTemplate(buf, configData, params.ParamMap); err != nil {
		log.WithFields(log.Fields{
			"template": params.Filename,
		}).Error("Error executing config template")
		logFileError(string(configData), err)
		return compiledTemplate, err
	}

	// parse the config yaml
	data := buf.Bytes()
	var config YamlConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		logFileError(string(data), err)
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

	// Process the core and plugin parsers
	conditions,
		metadata,
		mappings,
		outputs,
		parameters,
		resources,
		transform = processParsers(config.Resources, templateParsers)

	compiledTemplate = YamlCloudformation{
		AWSTemplateFormatVersion: config.AWSTemplateFormatVersion,
		Description:              config.Description,
		Metadata:                 mergeFinalTemplates(config.Metadata, metadata),
		Parameters:               mergeFinalTemplates(config.Parameters, parameters),
		Conditions:               mergeFinalTemplates(config.Conditions, conditions),
		Transform:                mergeFinalTemplates(config.Transform, transform),
		Mappings:                 mergeFinalTemplates(config.Mappings, mappings),
		Resources:                mergeFinalResources(config.Resources, resources),
		Outputs:                  mergeFinalTemplates(config.Outputs, outputs),
	}

	return compiledTemplate, nil
}

// Process the parser funcs against the template's resources
// and return new template objects
func processParsers(
	templateResources types.ResourceMap,
	parserFuncs map[string]types.ParserFunc,
) (
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
) {
	// Loop through each Resource in the template, and parse it with a ParserFunc
	for templateResourceName, templateResource := range templateResources {

		// If this is a custom resource, pass it through without touching it
		if templateResource.Type == "AWS::CloudFormation::CustomResource" ||
			strings.HasPrefix(templateResource.Type, "Custom::") {

			resources = mergeTemplates(
				templateResourceName,
				"aws-custom-resource",
				templateResource.Type,
				resources,
				types.TemplateObject{templateResourceName: templateResource},
			)
		} else { // This is a resource
			// Check if there is a parser for this resource
			parser, ok := parserFuncs[templateResource.Type]

			// If theres no parser log an error
			if !ok {
				printer.Error(
					fmt.Errorf("No parser found"),
					fmt.Sprintf(
						"\n   ├─ Name:    %s\n   ├─ Type:    %s\n   └─ Resolution:    %s",
						templateResourceName,
						templateResource.Type,
						"You may need to install a plugin to parse the resource.",
					),
					"",
				)
				continue
			}

			// Marshall the resource into YAML to send to the parser function
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
			conditions = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				conditions,
				parserConditions,
			)

			metadata = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				metadata,
				parserMetadata,
			)

			mappings = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				mappings,
				parserMappings,
			)

			outputs = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				outputs,
				parserOutputs,
			)

			parameters = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				parameters,
				parserParameters,
			)

			resources = mergeTemplates(
				templateResourceName,
				parserSource,
				templateResource.Type,
				resources,
				parserResources,
			)

			transform = mergeTemplates(
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

func logFileError(file string, err error) {
	errorLocation := -1
	re := regexp.MustCompile(`([0-9]+)`)
	match := re.FindStringSubmatch(err.Error())
	if len(match) > 1 {
		errorLocation, _ = strconv.Atoi(match[1])
	}

	lines := strings.Split(file, "\n")
	for nb, line := range lines {
		lineNb := nb + 1
		if lineNb == errorLocation {
			fmt.Printf(">>% 4d %v\n", lineNb, line)
		} else {
			fmt.Printf("% 6d %v\n", lineNb, line)
		}
	}
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

func mergeTemplates(
	name,
	source,
	resourceType string,
	maps ...map[string]interface{},
) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
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
	return result
}

func mergeFinalTemplates(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func mergeFinalResources(
	configResources types.ResourceMap,
	baseResources types.TemplateObject,
) (
	combinedResource types.TemplateObject,
) {

	for k, v := range configResources {
		if obj, err := json.Marshal(v); err == nil {
			var tempCfResource types.CfResource
			if err = json.Unmarshal(obj, &tempCfResource); err == nil {
				combinedResource[k] = tempCfResource
			}
		}
	}

	for k, v := range baseResources {
		if obj, err := json.Marshal(v); err == nil {
			var tempCfResource types.CfResource
			if err = json.Unmarshal(obj, &tempCfResource); err == nil {
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
