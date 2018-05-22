package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GameLiftBuild struct {
	Type       string                      `yaml:"Type"`
	Properties GameLiftBuildProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GameLiftBuildProperties struct {
	Name interface{} `yaml:"Name,omitempty"`
	Version interface{} `yaml:"Version,omitempty"`
	StorageLocation *properties.Build_S3Location `yaml:"StorageLocation,omitempty"`
}

func NewGameLiftBuild(properties GameLiftBuildProperties, deps ...interface{}) GameLiftBuild {
	return GameLiftBuild{
		Type:       "AWS::GameLift::Build",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGameLiftBuild(name string, data string) (cf types.ValueMap, err error) {
	var resource GameLiftBuild
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GameLiftBuild - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GameLiftBuild) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GameLiftBuildProperties) Validate() []error {
	errs := []error{}
	return errs
}
