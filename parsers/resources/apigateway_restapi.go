package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ApiGatewayRestApi struct {
	Type       string                      `yaml:"Type"`
	Properties ApiGatewayRestApiProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ApiGatewayRestApiProperties struct {
	ApiKeySourceType interface{} `yaml:"ApiKeySourceType,omitempty"`
	Body interface{} `yaml:"Body,omitempty"`
	CloneFrom interface{} `yaml:"CloneFrom,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	FailOnWarnings interface{} `yaml:"FailOnWarnings,omitempty"`
	MinimumCompressionSize interface{} `yaml:"MinimumCompressionSize,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	BodyS3Location *properties.RestApi_S3Location `yaml:"BodyS3Location,omitempty"`
	Parameters interface{} `yaml:"Parameters,omitempty"`
	BinaryMediaTypes interface{} `yaml:"BinaryMediaTypes,omitempty"`
	EndpointConfiguration *properties.RestApi_EndpointConfiguration `yaml:"EndpointConfiguration,omitempty"`
}

func NewApiGatewayRestApi(properties ApiGatewayRestApiProperties, deps ...interface{}) ApiGatewayRestApi {
	return ApiGatewayRestApi{
		Type:       "AWS::ApiGateway::RestApi",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseApiGatewayRestApi(name string, data string) (cf types.ValueMap, err error) {
	var resource ApiGatewayRestApi
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ApiGatewayRestApi - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ApiGatewayRestApi) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ApiGatewayRestApiProperties) Validate() []error {
	errs := []error{}
	return errs
}
