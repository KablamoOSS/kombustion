package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type CodePipelineCustomActionType struct {
	Type       string                      `yaml:"Type"`
	Properties CodePipelineCustomActionTypeProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type CodePipelineCustomActionTypeProperties struct {
	Category interface{} `yaml:"Category"`
	Provider interface{} `yaml:"Provider"`
	Version interface{} `yaml:"Version,omitempty"`
	Settings *properties.CustomActionType_Settings `yaml:"Settings,omitempty"`
	ConfigurationProperties interface{} `yaml:"ConfigurationProperties,omitempty"`
	InputArtifactDetails *properties.CustomActionType_ArtifactDetails `yaml:"InputArtifactDetails"`
	OutputArtifactDetails *properties.CustomActionType_ArtifactDetails `yaml:"OutputArtifactDetails"`
}

func NewCodePipelineCustomActionType(properties CodePipelineCustomActionTypeProperties, deps ...interface{}) CodePipelineCustomActionType {
	return CodePipelineCustomActionType{
		Type:       "AWS::CodePipeline::CustomActionType",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseCodePipelineCustomActionType(name string, data string) (cf types.ValueMap, err error) {
	var resource CodePipelineCustomActionType
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: CodePipelineCustomActionType - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource CodePipelineCustomActionType) Validate() []error {
	return resource.Properties.Validate()
}

func (resource CodePipelineCustomActionTypeProperties) Validate() []error {
	errs := []error{}
	if resource.Category == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Category'"))
	}
	if resource.Provider == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Provider'"))
	}
	if resource.InputArtifactDetails == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InputArtifactDetails'"))
	} else {
		errs = append(errs, resource.InputArtifactDetails.Validate()...)
	}
	if resource.OutputArtifactDetails == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OutputArtifactDetails'"))
	} else {
		errs = append(errs, resource.OutputArtifactDetails.Validate()...)
	}
	return errs
}
