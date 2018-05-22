package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type IoTThing struct {
	Type       string                      `yaml:"Type"`
	Properties IoTThingProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTThingProperties struct {
	ThingName interface{} `yaml:"ThingName,omitempty"`
	AttributePayload *properties.Thing_AttributePayload `yaml:"AttributePayload,omitempty"`
}

func NewIoTThing(properties IoTThingProperties, deps ...interface{}) IoTThing {
	return IoTThing{
		Type:       "AWS::IoT::Thing",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTThing(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTThing
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTThing - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTThing) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTThingProperties) Validate() []error {
	errs := []error{}
	return errs
}
