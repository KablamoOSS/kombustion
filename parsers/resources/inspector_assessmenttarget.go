package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type InspectorAssessmentTarget struct {
	Type       string                      `yaml:"Type"`
	Properties InspectorAssessmentTargetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type InspectorAssessmentTargetProperties struct {
	AssessmentTargetName interface{} `yaml:"AssessmentTargetName,omitempty"`
	ResourceGroupArn interface{} `yaml:"ResourceGroupArn"`
}

func NewInspectorAssessmentTarget(properties InspectorAssessmentTargetProperties, deps ...interface{}) InspectorAssessmentTarget {
	return InspectorAssessmentTarget{
		Type:       "AWS::Inspector::AssessmentTarget",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseInspectorAssessmentTarget(name string, data string) (cf types.ValueMap, err error) {
	var resource InspectorAssessmentTarget
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: InspectorAssessmentTarget - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource InspectorAssessmentTarget) Validate() []error {
	return resource.Properties.Validate()
}

func (resource InspectorAssessmentTargetProperties) Validate() []error {
	errs := []error{}
	if resource.ResourceGroupArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceGroupArn'"))
	}
	return errs
}
