package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GuardDutyFilter struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyFilterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyFilterProperties struct {
	Action interface{} `yaml:"Action"`
	Description interface{} `yaml:"Description"`
	DetectorId interface{} `yaml:"DetectorId"`
	Name interface{} `yaml:"Name,omitempty"`
	Rank interface{} `yaml:"Rank"`
	FindingCriteria *properties.Filter_FindingCriteria `yaml:"FindingCriteria"`
}

func NewGuardDutyFilter(properties GuardDutyFilterProperties, deps ...interface{}) GuardDutyFilter {
	return GuardDutyFilter{
		Type:       "AWS::GuardDuty::Filter",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyFilter(name string, data string) (cf types.ValueMap, err error) {
	var resource GuardDutyFilter
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyFilter - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GuardDutyFilter) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyFilterProperties) Validate() []error {
	errs := []error{}
	if resource.Action == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Action'"))
	}
	if resource.Description == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Description'"))
	}
	if resource.DetectorId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Rank == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Rank'"))
	}
	if resource.FindingCriteria == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FindingCriteria'"))
	} else {
		errs = append(errs, resource.FindingCriteria.Validate()...)
	}
	return errs
}
