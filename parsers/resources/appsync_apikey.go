package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type AppSyncApiKey struct {
	Type       string                      `yaml:"Type"`
	Properties AppSyncApiKeyProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type AppSyncApiKeyProperties struct {
	ApiId interface{} `yaml:"ApiId"`
	Description interface{} `yaml:"Description,omitempty"`
	Expires interface{} `yaml:"Expires,omitempty"`
}

func NewAppSyncApiKey(properties AppSyncApiKeyProperties, deps ...interface{}) AppSyncApiKey {
	return AppSyncApiKey{
		Type:       "AWS::AppSync::ApiKey",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseAppSyncApiKey(name string, data string) (cf types.ValueMap, err error) {
	var resource AppSyncApiKey
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: AppSyncApiKey - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource AppSyncApiKey) Validate() []error {
	return resource.Properties.Validate()
}

func (resource AppSyncApiKeyProperties) Validate() []error {
	errs := []error{}
	if resource.ApiId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApiId'"))
	}
	return errs
}
