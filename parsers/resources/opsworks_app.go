package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type OpsWorksApp struct {
	Type       string                      `yaml:"Type"`
	Properties OpsWorksAppProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type OpsWorksAppProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	EnableSsl interface{} `yaml:"EnableSsl,omitempty"`
	Name interface{} `yaml:"Name"`
	Shortname interface{} `yaml:"Shortname,omitempty"`
	StackId interface{} `yaml:"StackId"`
	Type interface{} `yaml:"Type"`
	SslConfiguration *properties.App_SslConfiguration `yaml:"SslConfiguration,omitempty"`
	AppSource *properties.App_Source `yaml:"AppSource,omitempty"`
	Attributes interface{} `yaml:"Attributes,omitempty"`
	DataSources interface{} `yaml:"DataSources,omitempty"`
	Domains interface{} `yaml:"Domains,omitempty"`
	Environment interface{} `yaml:"Environment,omitempty"`
}

func NewOpsWorksApp(properties OpsWorksAppProperties, deps ...interface{}) OpsWorksApp {
	return OpsWorksApp{
		Type:       "AWS::OpsWorks::App",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseOpsWorksApp(name string, data string) (cf types.ValueMap, err error) {
	var resource OpsWorksApp
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: OpsWorksApp - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource OpsWorksApp) Validate() []error {
	return resource.Properties.Validate()
}

func (resource OpsWorksAppProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.StackId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StackId'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
