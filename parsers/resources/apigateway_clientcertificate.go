package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ApiGatewayClientCertificate struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayClientCertificateProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayClientCertificateProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
}

func NewApiGatewayClientCertificate(properties ApiGatewayClientCertificateProperties, deps ...interface{}) ApiGatewayClientCertificate {
	return ApiGatewayClientCertificate{
		Type:       "AWS::ApiGateway::ClientCertificate",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayClientCertificate(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayClientCertificate
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayClientCertificate - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayClientCertificate) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayClientCertificateProperties) Validate() []error {
	errs := []error{}
	return errs
}
