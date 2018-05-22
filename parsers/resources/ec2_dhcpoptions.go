package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type EC2DHCPOptions struct {
	Type       string                      `yaml:"Type"`
	Properties EC2DHCPOptionsProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2DHCPOptionsProperties struct {
	DomainName interface{} `yaml:"DomainName,omitempty"`
	NetbiosNodeType interface{} `yaml:"NetbiosNodeType,omitempty"`
	DomainNameServers interface{} `yaml:"DomainNameServers,omitempty"`
	NetbiosNameServers interface{} `yaml:"NetbiosNameServers,omitempty"`
	NtpServers interface{} `yaml:"NtpServers,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewEC2DHCPOptions(properties EC2DHCPOptionsProperties, deps ...interface{}) EC2DHCPOptions {
	return EC2DHCPOptions{
		Type:       "AWS::EC2::DHCPOptions",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2DHCPOptions(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2DHCPOptions
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2DHCPOptions - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2DHCPOptions) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2DHCPOptionsProperties) Validate() []error {
	errs := []error{}
	return errs
}
