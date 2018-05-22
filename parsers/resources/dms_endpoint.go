package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type DMSEndpoint struct {
	Type       string                      `yaml:"Type"`
	Properties DMSEndpointProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSEndpointProperties struct {
	CertificateArn interface{} `yaml:"CertificateArn,omitempty"`
	DatabaseName interface{} `yaml:"DatabaseName,omitempty"`
	EndpointIdentifier interface{} `yaml:"EndpointIdentifier,omitempty"`
	EndpointType interface{} `yaml:"EndpointType"`
	EngineName interface{} `yaml:"EngineName"`
	ExtraConnectionAttributes interface{} `yaml:"ExtraConnectionAttributes,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	Password interface{} `yaml:"Password,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	ServerName interface{} `yaml:"ServerName,omitempty"`
	SslMode interface{} `yaml:"SslMode,omitempty"`
	Username interface{} `yaml:"Username,omitempty"`
	S3Settings *properties.Endpoint_S3Settings `yaml:"S3Settings,omitempty"`
	MongoDbSettings *properties.Endpoint_MongoDbSettings `yaml:"MongoDbSettings,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	DynamoDbSettings *properties.Endpoint_DynamoDbSettings `yaml:"DynamoDbSettings,omitempty"`
}

func NewDMSEndpoint(properties DMSEndpointProperties, deps ...interface{}) DMSEndpoint {
	return DMSEndpoint{
		Type:       "AWS::DMS::Endpoint",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSEndpoint(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSEndpoint
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSEndpoint - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSEndpoint) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSEndpointProperties) Validate() []error {
	errs := []error{}
	if resource.EndpointType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EndpointType'"))
	}
	if resource.EngineName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EngineName'"))
	}
	return errs
}
