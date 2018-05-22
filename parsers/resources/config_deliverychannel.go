package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ConfigDeliveryChannel struct {
	Type       string                      `yaml:"Type"`
	Properties ConfigDeliveryChannelProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ConfigDeliveryChannelProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	S3BucketName interface{} `yaml:"S3BucketName"`
	S3KeyPrefix interface{} `yaml:"S3KeyPrefix,omitempty"`
	SnsTopicARN interface{} `yaml:"SnsTopicARN,omitempty"`
	ConfigSnapshotDeliveryProperties *properties.DeliveryChannel_ConfigSnapshotDeliveryProperties `yaml:"ConfigSnapshotDeliveryProperties,omitempty"`
}

func NewConfigDeliveryChannel(properties ConfigDeliveryChannelProperties, deps ...interface{}) ConfigDeliveryChannel {
	return ConfigDeliveryChannel{
		Type:       "AWS::Config::DeliveryChannel",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseConfigDeliveryChannel(name string, data string) (cf types.ValueMap, err error) {
	var resource ConfigDeliveryChannel
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ConfigDeliveryChannel - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ConfigDeliveryChannel) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ConfigDeliveryChannelProperties) Validate() []error {
	errs := []error{}
	if resource.S3BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errs
}
