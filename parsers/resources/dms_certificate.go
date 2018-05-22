package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type DMSCertificate struct {
	Type       string                      `yaml:"Type"`
	Properties DMSCertificateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSCertificateProperties struct {
	CertificateIdentifier interface{} `yaml:"CertificateIdentifier,omitempty"`
	CertificatePem interface{} `yaml:"CertificatePem,omitempty"`
	CertificateWallet interface{} `yaml:"CertificateWallet,omitempty"`
}

func NewDMSCertificate(properties DMSCertificateProperties, deps ...interface{}) DMSCertificate {
	return DMSCertificate{
		Type:       "AWS::DMS::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSCertificate(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSCertificate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSCertificateProperties) Validate() []error {
	errs := []error{}
	return errs
}
