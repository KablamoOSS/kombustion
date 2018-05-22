package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type OpsWorksStack struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksStackProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksStackProperties struct {
	AgentVersion interface{} `yaml:"AgentVersion,omitempty"`
	ClonePermissions interface{} `yaml:"ClonePermissions,omitempty"`
	CustomJson interface{} `yaml:"CustomJson,omitempty"`
	DefaultAvailabilityZone interface{} `yaml:"DefaultAvailabilityZone,omitempty"`
	DefaultInstanceProfileArn interface{} `yaml:"DefaultInstanceProfileArn"`
	DefaultOs interface{} `yaml:"DefaultOs,omitempty"`
	DefaultRootDeviceType interface{} `yaml:"DefaultRootDeviceType,omitempty"`
	DefaultSshKeyName interface{} `yaml:"DefaultSshKeyName,omitempty"`
	DefaultSubnetId interface{} `yaml:"DefaultSubnetId,omitempty"`
	EcsClusterArn interface{} `yaml:"EcsClusterArn,omitempty"`
	HostnameTheme interface{} `yaml:"HostnameTheme,omitempty"`
	Name interface{} `yaml:"Name"`
	ServiceRoleArn interface{} `yaml:"ServiceRoleArn"`
	SourceStackId interface{} `yaml:"SourceStackId,omitempty"`
	UseCustomCookbooks interface{} `yaml:"UseCustomCookbooks,omitempty"`
	UseOpsworksSecurityGroups interface{} `yaml:"UseOpsworksSecurityGroups,omitempty"`
	VpcId interface{} `yaml:"VpcId,omitempty"`
	ConfigurationManager *properties.Stack_StackConfigurationManager `yaml:"ConfigurationManager,omitempty"`
	CustomCookbooksSource *properties.Stack_Source `yaml:"CustomCookbooksSource,omitempty"`
	Attributes interface{} `yaml:"Attributes,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	CloneAppIds interface{} `yaml:"CloneAppIds,omitempty"`
	ElasticIps interface{} `yaml:"ElasticIps,omitempty"`
	RdsDbInstances interface{} `yaml:"RdsDbInstances,omitempty"`
	ChefConfiguration *properties.Stack_ChefConfiguration `yaml:"ChefConfiguration,omitempty"`
}

func NewOpsWorksStack(properties OpsWorksStackProperties, deps ...interface{}) OpsWorksStack {
	return OpsWorksStack{
		Type:       "AWS::OpsWorks::Stack",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksStack(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksStack
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksStack - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksStack) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksStackProperties) Validate() []error {
	errs := []error{}
	if resource.DefaultInstanceProfileArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DefaultInstanceProfileArn'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ServiceRoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRoleArn'"))
	}
	return errs
}
