package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DataPipelinePipeline struct {
	Type       string                      `yaml:"Type"`
	Properties DataPipelinePipelineProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DataPipelinePipelineProperties struct {
	Activate interface{} `yaml:"Activate,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	ParameterObjects interface{} `yaml:"ParameterObjects"`
	ParameterValues interface{} `yaml:"ParameterValues,omitempty"`
	PipelineObjects interface{} `yaml:"PipelineObjects,omitempty"`
	PipelineTags interface{} `yaml:"PipelineTags,omitempty"`
}

func NewDataPipelinePipeline(properties DataPipelinePipelineProperties, deps ...interface{}) DataPipelinePipeline {
	return DataPipelinePipeline{
		Type:       "AWS::DataPipeline::Pipeline",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDataPipelinePipeline(name string, data string) (cf types.ValueMap, err error) {
	var resource DataPipelinePipeline
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DataPipelinePipeline - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DataPipelinePipeline) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DataPipelinePipelineProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ParameterObjects == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterObjects'"))
	}
	return errs
}
