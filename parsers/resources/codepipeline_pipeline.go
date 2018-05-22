package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CodePipelinePipeline struct {
	Type       string                      `yaml:"Type"`
	Properties CodePipelinePipelineProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodePipelinePipelineProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	RestartExecutionOnUpdate interface{} `yaml:"RestartExecutionOnUpdate,omitempty"`
	RoleArn interface{} `yaml:"RoleArn"`
	DisableInboundStageTransitions interface{} `yaml:"DisableInboundStageTransitions,omitempty"`
	Stages interface{} `yaml:"Stages"`
	ArtifactStore *properties.Pipeline_ArtifactStore `yaml:"ArtifactStore"`
}

func NewCodePipelinePipeline(properties CodePipelinePipelineProperties, deps ...interface{}) CodePipelinePipeline {
	return CodePipelinePipeline{
		Type:       "AWS::CodePipeline::Pipeline",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodePipelinePipeline(name string, data string) (cf types.ValueMap, err error) {
	var resource CodePipelinePipeline
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodePipelinePipeline - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodePipelinePipeline) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodePipelinePipelineProperties) Validate() []error {
	errs := []error{}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.Stages == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Stages'"))
	}
	if resource.ArtifactStore == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ArtifactStore'"))
	} else {
		errs = append(errs, resource.ArtifactStore.Validate()...)
	}
	return errs
}
