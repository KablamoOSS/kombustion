package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type ElasticsearchDomain struct {
	Type       string                      `yaml:"Type"`
	Properties ElasticsearchDomainProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ElasticsearchDomainProperties struct {
	AccessPolicies interface{} `yaml:"AccessPolicies,omitempty"`
	DomainName interface{} `yaml:"DomainName,omitempty"`
	ElasticsearchVersion interface{} `yaml:"ElasticsearchVersion,omitempty"`
	VPCOptions *properties.Domain_VPCOptions `yaml:"VPCOptions,omitempty"`
	SnapshotOptions *properties.Domain_SnapshotOptions `yaml:"SnapshotOptions,omitempty"`
	AdvancedOptions interface{} `yaml:"AdvancedOptions,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	ElasticsearchClusterConfig *properties.Domain_ElasticsearchClusterConfig `yaml:"ElasticsearchClusterConfig,omitempty"`
	EBSOptions *properties.Domain_EBSOptions `yaml:"EBSOptions,omitempty"`
}

func NewElasticsearchDomain(properties ElasticsearchDomainProperties, deps ...interface{}) ElasticsearchDomain {
	return ElasticsearchDomain{
		Type:       "AWS::Elasticsearch::Domain",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseElasticsearchDomain(name string, data string) (cf types.ValueMap, err error) {
	var resource ElasticsearchDomain
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ElasticsearchDomain - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ElasticsearchDomain) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ElasticsearchDomainProperties) Validate() []error {
	errs := []error{}
	return errs
}
