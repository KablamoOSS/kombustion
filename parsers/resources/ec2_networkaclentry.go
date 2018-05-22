package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EC2NetworkAclEntry struct {
	Type       string                      `yaml:"Type"`
	Properties EC2NetworkAclEntryProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2NetworkAclEntryProperties struct {
	CidrBlock interface{} `yaml:"CidrBlock"`
	Egress interface{} `yaml:"Egress,omitempty"`
	Ipv6CidrBlock interface{} `yaml:"Ipv6CidrBlock,omitempty"`
	NetworkAclId interface{} `yaml:"NetworkAclId"`
	Protocol interface{} `yaml:"Protocol"`
	RuleAction interface{} `yaml:"RuleAction"`
	RuleNumber interface{} `yaml:"RuleNumber"`
	PortRange *properties.NetworkAclEntry_PortRange `yaml:"PortRange,omitempty"`
	Icmp *properties.NetworkAclEntry_Icmp `yaml:"Icmp,omitempty"`
}

func NewEC2NetworkAclEntry(properties EC2NetworkAclEntryProperties, deps ...interface{}) EC2NetworkAclEntry {
	return EC2NetworkAclEntry{
		Type:       "AWS::EC2::NetworkAclEntry",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NetworkAclEntry(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2NetworkAclEntry
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NetworkAclEntry - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2NetworkAclEntry) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NetworkAclEntryProperties) Validate() []error {
	errs := []error{}
	if resource.CidrBlock == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CidrBlock'"))
	}
	if resource.NetworkAclId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NetworkAclId'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	if resource.RuleAction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleAction'"))
	}
	if resource.RuleNumber == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RuleNumber'"))
	}
	return errs
}
