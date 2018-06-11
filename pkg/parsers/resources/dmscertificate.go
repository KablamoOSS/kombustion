package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// DMSCertificate Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html
type DMSCertificate struct {
	Type       string                   `yaml:"Type"`
	Properties DMSCertificateProperties `yaml:"Properties"`
	Condition  interface{}              `yaml:"Condition,omitempty"`
	Metadata   interface{}              `yaml:"Metadata,omitempty"`
	DependsOn  interface{}              `yaml:"DependsOn,omitempty"`
}

// DMSCertificate Properties
type DMSCertificateProperties struct {
	CertificateIdentifier interface{} `yaml:"CertificateIdentifier,omitempty"`
	CertificatePem        interface{} `yaml:"CertificatePem,omitempty"`
	CertificateWallet     interface{} `yaml:"CertificateWallet,omitempty"`
}

// NewDMSCertificate constructor creates a new DMSCertificate
func NewDMSCertificate(properties DMSCertificateProperties, deps ...interface{}) DMSCertificate {
	return DMSCertificate{
		Type:       "AWS::DMS::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDMSCertificate parses DMSCertificate
func ParseDMSCertificate(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
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
	cf = types.TemplateObject{name: resource}
	return
}

func (resource DMSCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSCertificateProperties) Validate() []error {
	errs := []error{}
	return errs
}
