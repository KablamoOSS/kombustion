package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type GameLiftFleet struct {
	Type       string                      `yaml:"Type"`
	Properties GameLiftFleetProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GameLiftFleetProperties struct {
	BuildId interface{} `yaml:"BuildId"`
	Description interface{} `yaml:"Description,omitempty"`
	DesiredEC2Instances interface{} `yaml:"DesiredEC2Instances"`
	EC2InstanceType interface{} `yaml:"EC2InstanceType"`
	MaxSize interface{} `yaml:"MaxSize,omitempty"`
	MinSize interface{} `yaml:"MinSize,omitempty"`
	Name interface{} `yaml:"Name"`
	ServerLaunchParameters interface{} `yaml:"ServerLaunchParameters,omitempty"`
	ServerLaunchPath interface{} `yaml:"ServerLaunchPath"`
	EC2InboundPermissions interface{} `yaml:"EC2InboundPermissions,omitempty"`
	LogPaths interface{} `yaml:"LogPaths,omitempty"`
}

func NewGameLiftFleet(properties GameLiftFleetProperties, deps ...interface{}) GameLiftFleet {
	return GameLiftFleet{
		Type:       "AWS::GameLift::Fleet",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGameLiftFleet(name string, data string) (cf types.ValueMap, err error) {
	var resource GameLiftFleet
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GameLiftFleet - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GameLiftFleet) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GameLiftFleetProperties) Validate() []error {
	errs := []error{}
	if resource.BuildId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BuildId'"))
	}
	if resource.DesiredEC2Instances == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DesiredEC2Instances'"))
	}
	if resource.EC2InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EC2InstanceType'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ServerLaunchPath == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServerLaunchPath'"))
	}
	return errs
}
