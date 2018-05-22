package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApiGatewayDomainName struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayDomainNameProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayDomainNameProperties struct {
	CertificateArn interface{} `yaml:"CertificateArn,omitempty"`
	DomainName interface{} `yaml:"DomainName"`
	RegionalCertificateArn interface{} `yaml:"RegionalCertificateArn,omitempty"`
	EndpointConfiguration *properties.DomainName_EndpointConfiguration `yaml:"EndpointConfiguration,omitempty"`
}

func NewApiGatewayDomainName(properties ApiGatewayDomainNameProperties, deps ...interface{}) ApiGatewayDomainName {
	return ApiGatewayDomainName{
		Type:       "AWS::ApiGateway::DomainName",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayDomainName(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayDomainName
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayDomainName - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayDomainName) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayDomainNameProperties) Validate() []error {
	errs := []error{}
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	return errs
}
