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

	"github.com/KablamoOSS/kombustion/parsers"
	"github.com/KablamoOSS/kombustion/types"

	yaml "github.com/KablamoOSS/yaml"
)

type YamlConfig struct {
	AWSTemplateFormatVersion string            `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string            `yaml:"Description,omitempty"`
	Parameters               types.ValueMap    `yaml:"Parameters,omitempty"`
	Mappings                 types.ValueMap    `yaml:"Mappings,omitempty"`
	Conditions               types.ValueMap    `yaml:"Conditions,omitempty"`
	Transform                types.ValueMap    `yaml:"Transform,omitempty"`
	Resources                types.ResourceMap `yaml:"Resources"`
	Outputs                  types.ValueMap    `yaml:"Outputs,omitempty"`
}

// YamlCloudformation -
type YamlCloudformation struct {
	AWSTemplateFormatVersion string         `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string         `yaml:"Description,omitempty"`
	Parameters               types.ValueMap `yaml:"Parameters,omitempty"`
	Mappings                 types.ValueMap `yaml:"Mappings,omitempty"`
	Conditions               types.ValueMap `yaml:"Conditions,omitempty"`
	Transform                types.ValueMap `yaml:"Transform,omitempty"`
	Resources                types.ValueMap `yaml:"Resources"`
	Outputs                  types.ValueMap `yaml:"Outputs,omitempty"`
}

type GenerateParams struct {
	Filename           string
	EnvFile            string
	Env                string
	DisableBaseOutputs bool
	ParamMap           map[string]string
}

// ParserMap - a map of parsers
type ParserMap map[string]types.ParserFunc

var resourceParsers ParserMap
var outputParsers ParserMap
var mappingParsers ParserMap

func init() {
	registerYamlTagUnmarshalers()
}

func populateParsers(noBaseOutputs bool) {
	resourceParsers = parsers.GetParsers_resources()
	mappingParsers = make(ParserMap)

	if noBaseOutputs {
		outputParsers = make(ParserMap)
	} else {
		outputParsers = parsers.GetParsers_outputs()
	}

	r, o, m := loadPlugins()
	for k, v := range r {
		resourceParsers[k] = v
	}
	for k, v := range o {
		outputParsers[k] = v
	}
	for k, v := range m {
		mappingParsers[k] = v
	}
}

func PluginDocs() (docs map[string]string) {
	docs = make(map[string]string)
	r, _, _ := loadPlugins()
	for k := range r {
		// TODO: each plugin should export a `Usage` map.
		// this function should return those doc strings as values in the docs map
		docs[k] = ""
	}
	return
}

func DownloadPlugin(pluginname string) error {
	return downloadPlugin(pluginname)
}

func DeletePlugin(pluginname string) error {
	return deletePlugin(pluginname)
}

// GenerateYamlStack - generate a stack definition from ./configs
func GenerateYamlStack(params GenerateParams) (out YamlCloudformation, err error) {

	// load the config file
	var configData []byte

	// populate the parser variables
	populateParsers(params.DisableBaseOutputs)

	configPath := fmt.Sprintf(params.Filename)
	//configPath := fmt.Sprintf("./configs/%v.yaml", filename)
	if configData, err = ioutil.ReadFile(configPath); err != nil {
		return
	}

	// handle environment variables and custom params
	envMap := ResolveEnvironment(params.EnvFile, params.Env)
	for k, v := range params.ParamMap {
		envMap[k] = v
	}

	//preprocess - template in the environment variables and custom params
	buf := new(bytes.Buffer)
	if err = executeTemplate(buf, configData, envMap); err != nil {
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
	var outputs, resources, mappings types.ValueMap
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

func addBaseResources(baseResources types.ValueMap, configResources types.ResourceMap) (combinedResource types.ResourceMap) {
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

func yamlTemplateCF(resources types.ResourceMap, parsers ParserMap, warn bool) (compiled types.ValueMap, err error) {
	compiled = make(types.ValueMap)

	for resourceName, resource := range resources {
		if resource.Condition != nil { // if there is a condition on the source resource, warn the user
			log.WithFields(log.Fields{
				"resource": resourceName,
			}).Warn("Condition being applied on resource, this is not yet supported")
		}

		parser, ok := parsers[resource.Type]
		if !ok {
			if warn {
				log.WithFields(log.Fields{
					"type": resource.Type,
				}).Warn("Type not found")
			}
			continue
		}

		var resourseData []byte
		if resourseData, err = yaml.Marshal(resource); err != nil {
			return
		}

		var output types.ValueMap
		if output, err = parser(resourceName, string(resourseData)); err != nil {
			log.WithFields(log.Fields{
				"resource": resourceName,
			}).Error("Error parsing resource")
			logFileError(string(resourseData), err)
			return
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
