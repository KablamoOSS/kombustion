package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// EC2NetworkInterfaceAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface-attachment.html
type EC2NetworkInterfaceAttachment struct {
	Type       string                                  `yaml:"Type"`
	Properties EC2NetworkInterfaceAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                             `yaml:"Condition,omitempty"`
	Metadata   interface{}                             `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                             `yaml:"DependsOn,omitempty"`
}

// EC2NetworkInterfaceAttachment Properties
type EC2NetworkInterfaceAttachmentProperties struct {
	DeleteOnTermination interface{} `yaml:"DeleteOnTermination,omitempty"`
	DeviceIndex         interface{} `yaml:"DeviceIndex"`
	InstanceId          interface{} `yaml:"InstanceId"`
	NetworkInterfaceId  interface{} `yaml:"NetworkInterfaceId"`
}

// NewEC2NetworkInterfaceAttachment constructor creates a new EC2NetworkInterfaceAttachment
func NewEC2NetworkInterfaceAttachment(properties EC2NetworkInterfaceAttachmentProperties, deps ...interface{}) EC2NetworkInterfaceAttachment {
	return EC2NetworkInterfaceAttachment{
		Type:       "AWS::EC2::NetworkInterfaceAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2NetworkInterfaceAttachment parses EC2NetworkInterfaceAttachment
func ParseEC2NetworkInterfaceAttachment(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
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
