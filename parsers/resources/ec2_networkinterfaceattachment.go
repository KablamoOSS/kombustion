package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2NetworkInterfaceAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties EC2NetworkInterfaceAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2NetworkInterfaceAttachmentProperties struct {
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	DeviceIndex interface{} `yaml:"DeviceIndex"`
	InstanceId interface{} `yaml:"InstanceId"`
	NetworkInterfaceId interface{} `yaml:"NetworkInterfaceId"`
}

func NewEC2NetworkInterfaceAttachment(properties EC2NetworkInterfaceAttachmentProperties, deps ...interface{}) EC2NetworkInterfaceAttachment {
	return EC2NetworkInterfaceAttachment{
		Type:       "AWS::EC2::NetworkInterfaceAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2NetworkInterfaceAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2NetworkInterfaceAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2NetworkInterfaceAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2NetworkInterfaceAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2NetworkInterfaceAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.DeviceIndex == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DeviceIndex'"))
	}
	if resource.InstanceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceId'"))
	}
	if resource.NetworkInterfaceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NetworkInterfaceId'"))
	}
	return errs
}
