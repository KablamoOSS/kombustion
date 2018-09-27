package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// NeptuneDBCluster Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-neptune-dbcluster.html
type NeptuneDBCluster struct {
	Type       string                     `yaml:"Type"`
	Properties NeptuneDBClusterProperties `yaml:"Properties"`
	Condition  interface{}                `yaml:"Condition,omitempty"`
	Metadata   interface{}                `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                `yaml:"DependsOn,omitempty"`
}

// NeptuneDBCluster Properties
type NeptuneDBClusterProperties struct {
	BackupRetentionPeriod       interface{} `yaml:"BackupRetentionPeriod,omitempty"`
	DBClusterIdentifier         interface{} `yaml:"DBClusterIdentifier,omitempty"`
	DBClusterParameterGroupName interface{} `yaml:"DBClusterParameterGroupName,omitempty"`
	DBSubnetGroupName           interface{} `yaml:"DBSubnetGroupName,omitempty"`
	IamAuthEnabled              interface{} `yaml:"IamAuthEnabled,omitempty"`
	KmsKeyId                    interface{} `yaml:"KmsKeyId,omitempty"`
	Port                        interface{} `yaml:"Port,omitempty"`
	PreferredBackupWindow       interface{} `yaml:"PreferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow  interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	SnapshotIdentifier          interface{} `yaml:"SnapshotIdentifier,omitempty"`
	StorageEncrypted            interface{} `yaml:"StorageEncrypted,omitempty"`
	AvailabilityZones           interface{} `yaml:"AvailabilityZones,omitempty"`
	Tags                        interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds         interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

// NewNeptuneDBCluster constructor creates a new NeptuneDBCluster
func NewNeptuneDBCluster(properties NeptuneDBClusterProperties, deps ...interface{}) NeptuneDBCluster {
	return NeptuneDBCluster{
		Type:       "AWS::Neptune::DBCluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseNeptuneDBCluster parses NeptuneDBCluster
func ParseNeptuneDBCluster(
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
	var resource NeptuneDBCluster
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
					"Fn::Sub": "${AWS::StackName}-NeptuneDBCluster-" + name,
				},
			},
		},
	}

	return
}

// ParseNeptuneDBCluster validator
func (resource NeptuneDBCluster) Validate() []error {
	return resource.Properties.Validate()
}

// ParseNeptuneDBClusterProperties validator
func (resource NeptuneDBClusterProperties) Validate() []error {
	errors := []error{}
	return errors
}
