package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EC2EIPAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2EIPAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2EIPAssociationProperties struct {
	AllocationId interface{} `yaml:"AllocationId,omitempty"`
	EIP interface{} `yaml:"EIP,omitempty"`
	InstanceId interface{} `yaml:"InstanceId,omitempty"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress,omitempty"`
}

func NewEC2EIPAssociation(properties EC2EIPAssociationProperties, deps ...interface{}) EC2EIPAssociation {
	return EC2EIPAssociation{
		Type:       "AWS::EC2::EIPAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2EIPAssociation(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2EIPAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2EIPAssociation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2EIPAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2EIPAssociationProperties) Validate() []error {
	errs := []error{}
	return errs
}
