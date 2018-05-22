package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type DirectoryServiceSimpleAD struct {
	Type       string                      `yaml:"Type"`
	Properties DirectoryServiceSimpleADProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DirectoryServiceSimpleADProperties struct {
	CreateAlias interface{} `yaml:"CreateAlias,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	EnableSso interface{} `yaml:"EnableSso,omitempty"`
	Name interface{} `yaml:"Name"`
	Password interface{} `yaml:"Password"`
	ShortName interface{} `yaml:"ShortName,omitempty"`
	Size interface{} `yaml:"Size"`
	VpcSettings *properties.SimpleAD_VpcSettings `yaml:"VpcSettings"`
}

func NewDirectoryServiceSimpleAD(properties DirectoryServiceSimpleADProperties, deps ...interface{}) DirectoryServiceSimpleAD {
	return DirectoryServiceSimpleAD{
		Type:       "AWS::DirectoryService::SimpleAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDirectoryServiceSimpleAD(name string, data string) (cf types.ValueMap, err error) {
	var resource DirectoryServiceSimpleAD
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DirectoryServiceSimpleAD - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DirectoryServiceSimpleAD) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DirectoryServiceSimpleADProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.Size == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Size'"))
	}
	if resource.VpcSettings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcSettings'"))
	} else {
		errs = append(errs, resource.VpcSettings.Validate()...)
	}
	return errs
}
