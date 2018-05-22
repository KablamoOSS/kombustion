package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type GuardDutyIPSet struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyIPSetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyIPSetProperties struct {
	Activate interface{} `yaml:"Activate"`
	DetectorId interface{} `yaml:"DetectorId"`
	Format interface{} `yaml:"Format"`
	Location interface{} `yaml:"Location"`
	Name interface{} `yaml:"Name,omitempty"`
}

func NewGuardDutyIPSet(properties GuardDutyIPSetProperties, deps ...interface{}) GuardDutyIPSet {
	return GuardDutyIPSet{
		Type:       "AWS::GuardDuty::IPSet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyIPSet(name string, data string) (cf types.ValueMap, err error) {
	var resource GuardDutyIPSet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyIPSet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GuardDutyIPSet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyIPSetProperties) Validate() []error {
	errs := []error{}
	if resource.Activate == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Activate'"))
	}
	if resource.DetectorId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Format == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Format'"))
	}
	if resource.Location == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Location'"))
	}
	return errs
}
