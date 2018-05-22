package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type IoTCertificate struct {
	Type       string                      `yaml:"Type"`
	Properties IoTCertificateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type IoTCertificateProperties struct {
	CertificateSigningRequest interface{} `yaml:"CertificateSigningRequest"`
	Status interface{} `yaml:"Status"`
}

func NewIoTCertificate(properties IoTCertificateProperties, deps ...interface{}) IoTCertificate {
	return IoTCertificate{
		Type:       "AWS::IoT::Certificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseIoTCertificate(name string, data string) (cf types.ValueMap, err error) {
	var resource IoTCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: IoTCertificate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource IoTCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource IoTCertificateProperties) Validate() []error {
	errs := []error{}
	if resource.CertificateSigningRequest == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CertificateSigningRequest'"))
	}
	if resource.Status == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Status'"))
	}
	return errs
}
