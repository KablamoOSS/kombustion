package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AutoScalingLaunchConfiguration struct {
	Type       string                      `yaml:"Type"`
	Properties AutoScalingLaunchConfigurationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AutoScalingLaunchConfigurationProperties struct {
	AssociatePublicIpAddress interface{} `yaml:"AssociatePublicIpAddress,omitempty"`
	ClassicLinkVPCId interface{} `yaml:"ClassicLinkVPCId,omitempty"`
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	IamInstanceProfile interface{} `yaml:"IamInstanceProfile,omitempty"`
	ImageId interface{} `yaml:"ImageId"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	InstanceMonitoring interface{} `yaml:"InstanceMonitoring,omitempty"`
	InstanceType interface{} `yaml:"InstanceType"`
	KernelId interface{} `yaml:"KernelId,omitempty"`
	KeyName interface{} `yaml:"KeyName,omitempty"`
	PlacementTenancy interface{} `yaml:"PlacementTenancy,omitempty"`
	RamDiskId interface{} `yaml:"RamDiskId,omitempty"`
	SpotPrice interface{} `yaml:"SpotPrice,omitempty"`
	UserData interface{} `yaml:"UserData,omitempty"`
	BlockDeviceMappings interface{} `yaml:"BlockDeviceMappings,omitempty"`
	ClassicLinkVPCSecurityGroups interface{} `yaml:"ClassicLinkVPCSecurityGroups,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
}

func NewAutoScalingLaunchConfiguration(properties AutoScalingLaunchConfigurationProperties, deps ...interface{}) AutoScalingLaunchConfiguration {
	return AutoScalingLaunchConfiguration{
		Type:       "AWS::AutoScaling::LaunchConfiguration",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAutoScalingLaunchConfiguration(name string, data string) (cf types.ValueMap, err error) {
	var resource AutoScalingLaunchConfiguration
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AutoScalingLaunchConfiguration - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AutoScalingLaunchConfiguration) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AutoScalingLaunchConfigurationProperties) Validate() []error {
	errs := []error{}
	if resource.ImageId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ImageId'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
