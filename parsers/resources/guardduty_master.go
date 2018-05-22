package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type GuardDutyMaster struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyMasterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyMasterProperties struct {
	DetectorId interface{} `yaml:"DetectorId"`
	InvitationId interface{} `yaml:"InvitationId"`
	MasterId interface{} `yaml:"MasterId"`
}

func NewGuardDutyMaster(properties GuardDutyMasterProperties, deps ...interface{}) GuardDutyMaster {
	return GuardDutyMaster{
		Type:       "AWS::GuardDuty::Master",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyMaster(name string, data string) (cf types.ValueMap, err error) {
	var resource GuardDutyMaster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyMaster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GuardDutyMaster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyMasterProperties) Validate() []error {
	errs := []error{}
	if resource.DetectorId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.InvitationId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InvitationId'"))
	}
	if resource.MasterId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MasterId'"))
	}
	return errs
}
