package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2SecurityGroupIngress struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SecurityGroupIngressProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SecurityGroupIngressProperties struct {
	CidrIp interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6 interface{} `yaml:"CidrIpv6,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	FromPort interface{} `yaml:"FromPort,omitempty"`
	GroupId interface{} `yaml:"GroupId,omitempty"`
	GroupName interface{} `yaml:"GroupName,omitempty"`
	IpProtocol interface{} `yaml:"IpProtocol"`
	SourceSecurityGroupId interface{} `yaml:"SourceSecurityGroupId,omitempty"`
	SourceSecurityGroupName interface{} `yaml:"SourceSecurityGroupName,omitempty"`
	SourceSecurityGroupOwnerId interface{} `yaml:"SourceSecurityGroupOwnerId,omitempty"`
	ToPort interface{} `yaml:"ToPort,omitempty"`
}

func NewEC2SecurityGroupIngress(properties EC2SecurityGroupIngressProperties, deps ...interface{}) EC2SecurityGroupIngress {
	return EC2SecurityGroupIngress{
		Type:       "AWS::EC2::SecurityGroupIngress",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SecurityGroupIngress(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2SecurityGroupIngress
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SecurityGroupIngress - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2SecurityGroupIngress) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SecurityGroupIngressProperties) Validate() []error {
	errs := []error{}
	if resource.IpProtocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errs
}
