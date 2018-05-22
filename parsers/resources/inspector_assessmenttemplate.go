package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type InspectorAssessmentTemplate struct {
	Type       string                      `yaml:"Type"`
	Properties InspectorAssessmentTemplateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type InspectorAssessmentTemplateProperties struct {
	AssessmentTargetArn interface{} `yaml:"AssessmentTargetArn"`
	AssessmentTemplateName interface{} `yaml:"AssessmentTemplateName,omitempty"`
	DurationInSeconds interface{} `yaml:"DurationInSeconds"`
	RulesPackageArns interface{} `yaml:"RulesPackageArns"`
	UserAttributesForFindings interface{} `yaml:"UserAttributesForFindings,omitempty"`
}

func NewInspectorAssessmentTemplate(properties InspectorAssessmentTemplateProperties, deps ...interface{}) InspectorAssessmentTemplate {
	return InspectorAssessmentTemplate{
		Type:       "AWS::Inspector::AssessmentTemplate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseInspectorAssessmentTemplate(name string, data string) (cf types.ValueMap, err error) {
	var resource InspectorAssessmentTemplate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: InspectorAssessmentTemplate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource InspectorAssessmentTemplate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource InspectorAssessmentTemplateProperties) Validate() []error {
	errs := []error{}
	if resource.AssessmentTargetArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AssessmentTargetArn'"))
	}
	if resource.DurationInSeconds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DurationInSeconds'"))
	}
	if resource.RulesPackageArns == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RulesPackageArns'"))
	}
	return errs
}
