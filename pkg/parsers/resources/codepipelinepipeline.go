package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// CodePipelinePipeline Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
type CodePipelinePipeline struct {
	Type       string                         `yaml:"Type"`
	Properties CodePipelinePipelineProperties `yaml:"Properties"`
	Condition  interface{}                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                    `yaml:"DependsOn,omitempty"`
}

// CodePipelinePipeline Properties
type CodePipelinePipelineProperties struct {
	Name                           interface{}                       `yaml:"Name,omitempty"`
	RestartExecutionOnUpdate       interface{}                       `yaml:"RestartExecutionOnUpdate,omitempty"`
	RoleArn                        interface{}                       `yaml:"RoleArn"`
	DisableInboundStageTransitions interface{}                       `yaml:"DisableInboundStageTransitions,omitempty"`
	Stages                         interface{}                       `yaml:"Stages"`
	ArtifactStore                  *properties.PipelineArtifactStore `yaml:"ArtifactStore"`
}

// NewCodePipelinePipeline constructor creates a new CodePipelinePipeline
func NewCodePipelinePipeline(properties CodePipelinePipelineProperties, deps ...interface{}) CodePipelinePipeline {
	return CodePipelinePipeline{
		Type:       "AWS::CodePipeline::Pipeline",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCodePipelinePipeline parses CodePipelinePipeline
func ParseCodePipelinePipeline(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
