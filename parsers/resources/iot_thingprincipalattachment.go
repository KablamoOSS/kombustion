package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IoTThingPrincipalAttachment struct {
	Type       string                      `yaml:"Type"`
	Properties IoTThingPrincipalAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTThingPrincipalAttachmentProperties struct {
	Principal interface{} `yaml:"Principal"`
	ThingName interface{} `yaml:"ThingName"`
}

func NewIoTThingPrincipalAttachment(properties IoTThingPrincipalAttachmentProperties, deps ...interface{}) IoTThingPrincipalAttachment {
	return IoTThingPrincipalAttachment{
		Type:       "AWS::IoT::ThingPrincipalAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTThingPrincipalAttachment(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTThingPrincipalAttachment
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTThingPrincipalAttachment - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTThingPrincipalAttachment) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTThingPrincipalAttachmentProperties) Validate() []error {
	errs := []error{}
	if resource.Principal == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Principal'"))
	}
	if resource.ThingName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ThingName'"))
	}
	return errs
}
