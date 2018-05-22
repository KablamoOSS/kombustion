package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type GlueCrawler struct {
	Type       string                      `yaml:"Type"`
	Properties GlueCrawlerProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type GlueCrawlerProperties struct {
	DatabaseName interface{} `yaml:"DatabaseName"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	Role interface{} `yaml:"Role"`
	TablePrefix interface{} `yaml:"TablePrefix,omitempty"`
	Targets *properties.Crawler_Targets `yaml:"Targets"`
	SchemaChangePolicy *properties.Crawler_SchemaChangePolicy `yaml:"SchemaChangePolicy,omitempty"`
	Schedule *properties.Crawler_Schedule `yaml:"Schedule,omitempty"`
	Classifiers interface{} `yaml:"Classifiers,omitempty"`
}

func NewGlueCrawler(properties GlueCrawlerProperties, deps ...interface{}) GlueCrawler {
	return GlueCrawler{
		Type:       "AWS::Glue::Crawler",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseGlueCrawler(name string, data string) (cf types.ValueMap, err error) {
	var resource GlueCrawler
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: GlueCrawler - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource GlueCrawler) Validate() []error {
	return resource.Properties.Validate()
}

func (resource GlueCrawlerProperties) Validate() []error {
	errs := []error{}
	if resource.DatabaseName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DatabaseName'"))
	}
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Targets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Targets'"))
	} else {
		errs = append(errs, resource.Targets.Validate()...)
	}
	return errs
}
