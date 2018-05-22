package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GameLiftAlias struct {
	Type       string                      `yaml:"Type"`
	Properties GameLiftAliasProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GameLiftAliasProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	RoutingStrategy *properties.Alias_RoutingStrategy `yaml:"RoutingStrategy"`
}

func NewGameLiftAlias(properties GameLiftAliasProperties, deps ...interface{}) GameLiftAlias {
	return GameLiftAlias{
		Type:       "AWS::GameLift::Alias",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGameLiftAlias(name string, data string) (cf types.ValueMap, err error) {
	var resource GameLiftAlias
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GameLiftAlias - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GameLiftAlias) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GameLiftAliasProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.RoutingStrategy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoutingStrategy'"))
	} else {
		errs = append(errs, resource.RoutingStrategy.Validate()...)
	}
	return errs
}
