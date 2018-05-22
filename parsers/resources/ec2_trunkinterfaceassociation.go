package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2TrunkInterfaceAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2TrunkInterfaceAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2TrunkInterfaceAssociationProperties struct {
	BranchInterfaceId interface{} `yaml:"BranchInterfaceId"`
	GREKey interface{} `yaml:"GREKey,omitempty"`
	TrunkInterfaceId interface{} `yaml:"TrunkInterfaceId"`
	VLANId interface{} `yaml:"VLANId,omitempty"`
}

func NewEC2TrunkInterfaceAssociation(properties EC2TrunkInterfaceAssociationProperties, deps ...interface{}) EC2TrunkInterfaceAssociation {
	return EC2TrunkInterfaceAssociation{
		Type:       "AWS::EC2::TrunkInterfaceAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2TrunkInterfaceAssociation(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2TrunkInterfaceAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2TrunkInterfaceAssociation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2TrunkInterfaceAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2TrunkInterfaceAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.BranchInterfaceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BranchInterfaceId'"))
	}
	if resource.TrunkInterfaceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TrunkInterfaceId'"))
	}
	return errs
}
