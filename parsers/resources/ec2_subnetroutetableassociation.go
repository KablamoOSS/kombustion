package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2SubnetRouteTableAssociation struct {
	Type       string                      `yaml:"Type"`
	Properties EC2SubnetRouteTableAssociationProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2SubnetRouteTableAssociationProperties struct {
	RouteTableId interface{} `yaml:"RouteTableId"`
	SubnetId interface{} `yaml:"SubnetId"`
}

func NewEC2SubnetRouteTableAssociation(properties EC2SubnetRouteTableAssociationProperties, deps ...interface{}) EC2SubnetRouteTableAssociation {
	return EC2SubnetRouteTableAssociation{
		Type:       "AWS::EC2::SubnetRouteTableAssociation",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2SubnetRouteTableAssociation(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2SubnetRouteTableAssociation
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2SubnetRouteTableAssociation - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2SubnetRouteTableAssociation) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2SubnetRouteTableAssociationProperties) Validate() []error {
	errs := []error{}
	if resource.RouteTableId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RouteTableId'"))
	}
	if resource.SubnetId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetId'"))
	}
	return errs
}
