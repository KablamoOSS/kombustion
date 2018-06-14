package cloudformation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/pkg/parsers"
	"github.com/KablamoOSS/kombustion/types"

	yaml "github.com/KablamoOSS/yaml"
)

// YamlConfig -
type YamlConfig struct {
	AWSTemplateFormatVersion string               `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string               `yaml:"Description,omitempty"`
	Parameters               types.TemplateObject `yaml:"Parameters,omitempty"`
	Mappings                 types.TemplateObject `yaml:"Mappings,omitempty"`
	Conditions               types.TemplateObject `yaml:"Conditions,omitempty"`
	Transform                types.TemplateObject `yaml:"Transform,omitempty"`
	Resources                types.ResourceMap    `yaml:"Resources"`
	Outputs                  types.TemplateObject `yaml:"Outputs,omitempty"`
}

// YamlCloudformation -
type YamlCloudformation struct {
	AWSTemplateFormatVersion string               `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string               `yaml:"Description,omitempty"`
	Parameters               types.TemplateObject `yaml:"Parameters,omitempty"`
	Mappings                 types.TemplateObject `yaml:"Mappings,omitempty"`
	Conditions               types.TemplateObject `yaml:"Conditions,omitempty"`
	Transform                types.TemplateObject `yaml:"Transform,omitempty"`
	Resources                types.TemplateObject `yaml:"Resources"`
	Outputs                  types.TemplateObject `yaml:"Outputs,omitempty"`
}

// GenerateParams are required to generate a cloudformation yaml template
type GenerateParams struct {
	Filename           string
	EnvFile            string
	Env                string
	DisableBaseOutputs bool
	ParamMap           map[string]string
	Plugins            []*plugins.PluginLoaded
}

var resourceParsers map[string]types.ParserFunc
var outputParsers map[string]types.ParserFunc
var mappingParsers map[string]types.ParserFunc

func init() {
	registerYamlTagUnmarshalers()
}

func populateParsers(loadedPlugins []*plugins.PluginLoaded, noBaseOutputs bool) {
	resourceParsers = parsers.GetParsersResources()
	mappingParsers = make(map[string]types.ParserFunc)

	if noBaseOutputs {
		outputParsers = make(map[string]types.ParserFunc)
	} else {
		outputParsers = parsers.GetParsersOutputs()
	}

	plugins.ExtractResourcesFromPlugins(loadedPlugins, &resourceParsers)
	plugins.ExtractMappingsFromPlugins(loadedPlugins, &mappingParsers)
	plugins.ExtractOutputsFromPlugins(loadedPlugins, &outputParsers)

	// dont need if above extracts work
	// resources, outputs, mappings := plugins.LoadPlugins()
	// for k, v := range resources {
	// 	resourceParsers[k] = v
	// }
	// for k, v := range outputs {
	// 	outputParsers[k] = v
	// }
	// for k, v := range mappings {
	// 	mappingParsers[k] = v
	// }
}

// PluginDocs -
func PluginDocs() (docs map[string]string) {
	docs = make(map[string]string)
	// TODO: Plugins need to be passed in now
	// r, _, _ := plugins.LoadPlugins()
	// for k := range r {
	// TODO: each plugin should export a `Usage` map.
	// this function should return those doc strings as values in the docs map
	// docs[k] = ""
	// }
	return
}

// GenerateYamlStack - generate a stack definition from ./configs
func GenerateYamlStack(params GenerateParams) (out YamlCloudformation, err error) {

	// load the config file
	var configData []byte

	// populate the parser variables
	populateParsers(params.Plugins, params.DisableBaseOutputs)

	configPath := fmt.Sprintf(params.Filename)
	//configPath := fmt.Sprintf("./configs/%v.yaml", filename)
	if configData, err = ioutil.ReadFile(configPath); err != nil {
		return
	}

	//preprocess - template in the environment variables and custom params
	buf := new(bytes.Buffer)
	if err = executeTemplate(buf, configData, params.ParamMap); err != nil {
		log.WithFields(log.Fields{
			"template": configPath,
		}).Error("Error executing config template")
		logFileError(string(configData), err)
		return
	}

	// parse the config yaml
	data := buf.Bytes()
	var config YamlConfig
	if err = yaml.Unmarshal(data, &config); err != nil {
		logFileError(string(data), err)
		return
	}

	// compile the cloudformation
	var outputs, resources, mappings types.TemplateObject
	if resources, err = yamlTemplateCF(config.Resources, resourceParsers, true); err != nil {
		return
	}

	//Adding(Replacing) base objects for correct outputs by type
	config.Resources = addBaseResources(resources, config.Resources)

	if outputs, err = yamlTemplateCF(config.Resources, outputParsers, false); err != nil {
		return
	}
	if mappings, err = yamlTemplateCF(config.Resources, mappingParsers, false); err != nil {
		return
	}

	// merge mappings
	for k, v := range config.Mappings {
		if _, ok := mappings[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"mapping": k,
			}).Warn("duplicate mapping definition for mapping - overwriting")
		}
		mappings[k] = v
	}

	// merge outputs
	for k, v := range config.Outputs {
		if _, ok := outputs[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"output": k,
			}).Warn("duplicate output definition for output - overwriting")
		}
		outputs[k] = v
	}

	out = YamlCloudformation{
		AWSTemplateFormatVersion: config.AWSTemplateFormatVersion,
		Description:              config.Description,
		Parameters:               config.Parameters,
		Conditions:               config.Conditions,
		Transform:                config.Transform,
		Mappings:                 mappings,
		Resources:                resources,
		Outputs:                  outputs,
	}

	return
}

func addBaseResources(baseResources types.TemplateObject, configResources types.ResourceMap) (combinedResource types.ResourceMap) {
	combinedResource = configResources
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

func yamlTemplateCF(
	resources types.ResourceMap,
	parsers map[string]types.ParserFunc,
	isResources bool,
) (
	compiled types.TemplateObject,
	err error,
) {
	compiled = make(types.TemplateObject)

	for resourceName, resource := range resources {
		if resource.Condition != nil { // if there is a condition on the source resource, warn the user
			log.WithFields(log.Fields{
				"resource": resourceName,
			}).Warn("Condition being applied on resource, this is not yet supported")
		}

		var output types.TemplateObject
		var resourseData []byte

		if isResources && (resource.Type == "AWS::CloudFormation::CustomResource" || strings.HasPrefix(resource.Type, "Custom::")) {
			var cfResource types.CfResource

			if resourseData, err = yaml.Marshal(resource); err != nil {
				return
			}

			if err = yaml.Unmarshal([]byte(resourseData), &cfResource); err != nil {
				return
			}

			output = types.TemplateObject{resourceName: cfResource}
		} else {
			parser, ok := parsers[resource.Type]
			if !ok {
				if isResources {
					log.WithFields(log.Fields{
						"type": resource.Type,
					}).Warn("Type not found")
				}
				continue
			}

			if resourseData, err = yaml.Marshal(resource); err != nil {
				return
			}

			if output, err = parser(resourceName, string(resourseData)); err != nil {
				log.WithFields(log.Fields{
					"resource": resourceName,
				}).Error("Error parsing resource")
				logFileError(string(resourseData), err)
				return
			}
		}

		// collect all output resources in one list
		for k, v := range output {
			compiled[k] = v
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
