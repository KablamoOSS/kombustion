package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type EC2Volume struct {
	Type       string                      `yaml:"Type"`
	Properties EC2VolumeProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2VolumeProperties struct {
	AutoEnableIO interface{} `yaml:"AutoEnableIO,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	Size interface{} `yaml:"Size,omitempty"`
	SnapshotId interface{} `yaml:"SnapshotId,omitempty"`
	VolumeType interface{} `yaml:"VolumeType,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2Volume(properties EC2VolumeProperties, deps ...interface{}) EC2Volume {
	return EC2Volume{
		Type:       "AWS::EC2::Volume",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2Volume(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2Volume
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Volume - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2Volume) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2VolumeProperties) Validate() []error {
	errs := []error{}
	if resource.AvailabilityZone == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AvailabilityZone'"))
	}
	return errs
}
