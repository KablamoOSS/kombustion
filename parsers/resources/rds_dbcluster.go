package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBCluster struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBClusterProperties struct {
	BackupRetentionPeriod interface{} `yaml:"BackupRetentionPeriod,omitempty"`
	DBClusterIdentifier interface{} `yaml:"DBClusterIdentifier,omitempty"`
	DBClusterParameterGroupName interface{} `yaml:"DBClusterParameterGroupName,omitempty"`
	DBSubnetGroupName interface{} `yaml:"DBSubnetGroupName,omitempty"`
	DatabaseName interface{} `yaml:"DatabaseName,omitempty"`
	Engine interface{} `yaml:"Engine"`
	EngineVersion interface{} `yaml:"EngineVersion,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	MasterUserPassword interface{} `yaml:"MasterUserPassword,omitempty"`
	MasterUsername interface{} `yaml:"MasterUsername,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	PreferredBackupWindow interface{} `yaml:"PreferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	ReplicationSourceIdentifier interface{} `yaml:"ReplicationSourceIdentifier,omitempty"`
	SnapshotIdentifier interface{} `yaml:"SnapshotIdentifier,omitempty"`
	StorageEncrypted interface{} `yaml:"StorageEncrypted,omitempty"`
	AvailabilityZones interface{} `yaml:"AvailabilityZones,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

func NewRDSDBCluster(properties RDSDBClusterProperties, deps ...interface{}) RDSDBCluster {
	return RDSDBCluster{
		Type:       "AWS::RDS::DBCluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBClusterProperties) Validate() []error {
	errs := []error{}
	if resource.Engine == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Engine'"))
	}
	return errs
}
