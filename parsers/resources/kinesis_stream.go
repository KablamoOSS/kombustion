package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type KinesisStream struct {
	Type       string                      `yaml:"Type"`
	Properties KinesisStreamProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type KinesisStreamProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	RetentionPeriodHours interface{} `yaml:"RetentionPeriodHours,omitempty"`
	ShardCount interface{} `yaml:"ShardCount"`
	StreamEncryption *properties.Stream_StreamEncryption `yaml:"StreamEncryption,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewKinesisStream(properties KinesisStreamProperties, deps ...interface{}) KinesisStream {
	return KinesisStream{
		Type:       "AWS::Kinesis::Stream",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseKinesisStream(name string, data string) (cf types.ValueMap, err error) {
	var resource KinesisStream
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: KinesisStream - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource KinesisStream) Validate() []error {
	return resource.Properties.Validate()
}

func (resource KinesisStreamProperties) Validate() []error {
	errs := []error{}
	if resource.ShardCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ShardCount'"))
	}
	return errs
}
