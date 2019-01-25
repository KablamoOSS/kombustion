package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EMRCluster Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-cluster.html
type EMRCluster struct {
	Type       string               `yaml:"Type"`
	Properties EMRClusterProperties `yaml:"Properties"`
	Condition  interface{}          `yaml:"Condition,omitempty"`
	Metadata   interface{}          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}          `yaml:"DependsOn,omitempty"`
}

// EMRCluster Properties
type EMRClusterProperties struct {
	AdditionalInfo        interface{} `yaml:"AdditionalInfo,omitempty"`
	AutoScalingRole       interface{} `yaml:"AutoScalingRole,omitempty"`
	CustomAmiId           interface{} `yaml:"CustomAmiId,omitempty"`
	EbsRootVolumeSize     interface{} `yaml:"EbsRootVolumeSize,omitempty"`
	JobFlowRole           interface{} `yaml:"JobFlowRole"`
	LogUri                interface{} `yaml:"LogUri,omitempty"`
	Name                  interface{} `yaml:"Name"`
	ReleaseLabel          interface{} `yaml:"ReleaseLabel,omitempty"`
	ScaleDownBehavior     interface{} `yaml:"ScaleDownBehavior,omitempty"`
	SecurityConfiguration interface{} `yaml:"SecurityConfiguration,omitempty"`
	ServiceRole           interface{} `yaml:"ServiceRole"`
	VisibleToAllUsers     interface{} `yaml:"VisibleToAllUsers,omitempty"`
	Applications          interface{} `yaml:"Applications,omitempty"`
	BootstrapActions      interface{} `yaml:"BootstrapActions,omitempty"`
	Configurations        interface{} `yaml:"Configurations,omitempty"`
	Steps                 interface{} `yaml:"Steps,omitempty"`
	Tags                  interface{} `yaml:"Tags,omitempty"`
	KerberosAttributes    interface{} `yaml:"KerberosAttributes,omitempty"`
	Instances             interface{} `yaml:"Instances"`
}

// NewEMRCluster constructor creates a new EMRCluster
func NewEMRCluster(properties EMRClusterProperties, deps ...interface{}) EMRCluster {
	return EMRCluster{
		Type:       "AWS::EMR::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEMRCluster parses EMRCluster
func ParseEMRCluster(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource EMRCluster
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EMRCluster-" + name,
				},
			},
		},
	}

	return
}

// ParseEMRCluster validator
func (resource EMRCluster) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEMRClusterProperties validator
func (resource EMRClusterProperties) Validate() []error {
	errors := []error{}
	return errors
}
