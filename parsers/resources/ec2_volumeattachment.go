package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2VolumeAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VolumeAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VolumeAttachmentProperties struct {
	Device interface{} `yaml:"Device"`
	InstanceId interface{} `yaml:"InstanceId"`
	VolumeId interface{} `yaml:"VolumeId"`
}

func NewEC2VolumeAttachment(properties EC2VolumeAttachmentProperties, deps ...interface{}) EC2VolumeAttachment {
	return EC2VolumeAttachment{
		Type:       "AWS::EC2::VolumeAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2VolumeAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2VolumeAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2VolumeAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2VolumeAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VolumeAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.Device == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Device'"))
	}
	if resource.InstanceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceId'"))
	}
	if resource.VolumeId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VolumeId'"))
	}
	return errs
}
