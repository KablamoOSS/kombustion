package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type OpsWorksInstance struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksInstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksInstanceProperties struct {
	AgentVersion interface{} `yaml:"AgentVersion,omitempty"`
	AmiId interface{} `yaml:"AmiId,omitempty"`
	Architecture interface{} `yaml:"Architecture,omitempty"`
	AutoScalingType interface{} `yaml:"AutoScalingType,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	Hostname interface{} `yaml:"Hostname,omitempty"`
	InstallUpdatesOnBoot interface{} `yaml:"InstallUpdatesOnBoot,omitempty"`
	InstanceType interface{} `yaml:"InstanceType"`
	Os interface{} `yaml:"Os,omitempty"`
	RootDeviceType interface{} `yaml:"RootDeviceType,omitempty"`
	SshKeyName interface{} `yaml:"SshKeyName,omitempty"`
	StackId interface{} `yaml:"StackId"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	Tenancy interface{} `yaml:"Tenancy,omitempty"`
	VirtualizationType interface{} `yaml:"VirtualizationType,omitempty"`
	TimeBasedAutoScaling *properties.Instance_TimeBasedAutoScaling `yaml:"TimeBasedAutoScaling,omitempty"`
	BlockDeviceMappings interface{} `yaml:"BlockDeviceMappings,omitempty"`
	ElasticIps interface{} `yaml:"ElasticIps,omitempty"`
	LayerIds interface{} `yaml:"LayerIds"`
	Volumes interface{} `yaml:"Volumes,omitempty"`
}

func NewOpsWorksInstance(properties OpsWorksInstanceProperties, deps ...interface{}) OpsWorksInstance {
	return OpsWorksInstance{
		Type:       "AWS::OpsWorks::Instance",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksInstance(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksInstance - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksInstance) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	if resource.StackId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StackId'"))
	}
	if resource.LayerIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LayerIds'"))
	}
	return errs
}
