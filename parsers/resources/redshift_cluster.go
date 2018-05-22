package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type RedshiftCluster struct {
	Type       string                      `yaml:"Type"`
	Properties RedshiftClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RedshiftClusterProperties struct {
	AllowVersionUpgrade interface{} `yaml:"AllowVersionUpgrade,omitempty"`
	AutomatedSnapshotRetentionPeriod interface{} `yaml:"AutomatedSnapshotRetentionPeriod,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	ClusterIdentifier interface{} `yaml:"ClusterIdentifier,omitempty"`
	ClusterParameterGroupName interface{} `yaml:"ClusterParameterGroupName,omitempty"`
	ClusterSubnetGroupName interface{} `yaml:"ClusterSubnetGroupName,omitempty"`
	ClusterType interface{} `yaml:"ClusterType"`
	ClusterVersion interface{} `yaml:"ClusterVersion,omitempty"`
	DBName interface{} `yaml:"DBName"`
	ElasticIp interface{} `yaml:"ElasticIp,omitempty"`
	Encrypted interface{} `yaml:"Encrypted,omitempty"`
	HsmClientCertificateIdentifier interface{} `yaml:"HsmClientCertificateIdentifier,omitempty"`
	HsmConfigurationIdentifier interface{} `yaml:"HsmConfigurationIdentifier,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	MasterUserPassword interface{} `yaml:"MasterUserPassword"`
	MasterUsername interface{} `yaml:"MasterUsername"`
	NodeType interface{} `yaml:"NodeType"`
	NumberOfNodes interface{} `yaml:"NumberOfNodes,omitempty"`
	OwnerAccount interface{} `yaml:"OwnerAccount,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	PubliclyAccessible interface{} `yaml:"PubliclyAccessible,omitempty"`
	SnapshotClusterIdentifier interface{} `yaml:"SnapshotClusterIdentifier,omitempty"`
	SnapshotIdentifier interface{} `yaml:"SnapshotIdentifier,omitempty"`
	LoggingProperties *properties.Cluster_LoggingProperties `yaml:"LoggingProperties,omitempty"`
	ClusterSecurityGroups interface{} `yaml:"ClusterSecurityGroups,omitempty"`
	IamRoles interface{} `yaml:"IamRoles,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

func NewRedshiftCluster(properties RedshiftClusterProperties, deps ...interface{}) RedshiftCluster {
	return RedshiftCluster{
		Type:       "AWS::Redshift::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRedshiftCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource RedshiftCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RedshiftCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RedshiftCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RedshiftClusterProperties) Validate() []error {
	errs := []error{}
	if resource.ClusterType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ClusterType'"))
	}
	if resource.DBName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBName'"))
	}
	if resource.MasterUserPassword == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MasterUserPassword'"))
	}
	if resource.MasterUsername == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MasterUsername'"))
	}
	if resource.NodeType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NodeType'"))
	}
	return errs
}
