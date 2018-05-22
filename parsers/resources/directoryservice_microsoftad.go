package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type DirectoryServiceMicrosoftAD struct {
	Type       string                      `yaml:"Type"`
	Properties DirectoryServiceMicrosoftADProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DirectoryServiceMicrosoftADProperties struct {
	CreateAlias interface{} `yaml:"CreateAlias,omitempty"`
	EnableSso interface{} `yaml:"EnableSso,omitempty"`
	Name interface{} `yaml:"Name"`
	Password interface{} `yaml:"Password"`
	ShortName interface{} `yaml:"ShortName,omitempty"`
	VpcSettings *properties.MicrosoftAD_VpcSettings `yaml:"VpcSettings"`
}

func NewDirectoryServiceMicrosoftAD(properties DirectoryServiceMicrosoftADProperties, deps ...interface{}) DirectoryServiceMicrosoftAD {
	return DirectoryServiceMicrosoftAD{
		Type:       "AWS::DirectoryService::MicrosoftAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDirectoryServiceMicrosoftAD(name string, data string) (cf types.ValueMap, err error) {
	var resource DirectoryServiceMicrosoftAD
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DirectoryServiceMicrosoftAD - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DirectoryServiceMicrosoftAD) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DirectoryServiceMicrosoftADProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.VpcSettings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcSettings'"))
	} else {
		errs = append(errs, resource.VpcSettings.Validate()...)
	}
	return errs
}
