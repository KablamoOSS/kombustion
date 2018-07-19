package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// DomainElasticsearchClusterConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html
type DomainElasticsearchClusterConfig struct {
	DedicatedMasterCount   interface{} `yaml:"DedicatedMasterCount,omitempty"`
	DedicatedMasterEnabled interface{} `yaml:"DedicatedMasterEnabled,omitempty"`
	DedicatedMasterType    interface{} `yaml:"DedicatedMasterType,omitempty"`
	InstanceCount          interface{} `yaml:"InstanceCount,omitempty"`
	InstanceType           interface{} `yaml:"InstanceType,omitempty"`
	ZoneAwarenessEnabled   interface{} `yaml:"ZoneAwarenessEnabled,omitempty"`
}

// DomainElasticsearchClusterConfig validation
func (resource DomainElasticsearchClusterConfig) Validate() []error {
	errors := []error{}

	return errors
}
