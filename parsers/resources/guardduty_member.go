package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type GuardDutyMember struct {
	Type       string                      `yaml:"Type"`
	Properties GuardDutyMemberProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GuardDutyMemberProperties struct {
	DetectorId interface{} `yaml:"DetectorId"`
	DisableEmailNotification interface{} `yaml:"DisableEmailNotification,omitempty"`
	Email interface{} `yaml:"Email"`
	MemberId interface{} `yaml:"MemberId"`
	Message interface{} `yaml:"Message,omitempty"`
	Status interface{} `yaml:"Status,omitempty"`
}

func NewGuardDutyMember(properties GuardDutyMemberProperties, deps ...interface{}) GuardDutyMember {
	return GuardDutyMember{
		Type:       "AWS::GuardDuty::Member",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGuardDutyMember(name string, data string) (cf types.ValueMap, err error) {
	var resource GuardDutyMember
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GuardDutyMember - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GuardDutyMember) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GuardDutyMemberProperties) Validate() []error {
	errs := []error{}
	if resource.DetectorId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DetectorId'"))
	}
	if resource.Email == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Email'"))
	}
	if resource.MemberId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MemberId'"))
	}
	return errs
}
