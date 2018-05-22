package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type RDSDBInstance struct {
	Type       string                      `yaml:"Type"`
	Properties RDSDBInstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type RDSDBInstanceProperties struct {
	AllocatedStorage interface{} `yaml:"AllocatedStorage,omitempty"`
	AllowMajorVersionUpgrade interface{} `yaml:"AllowMajorVersionUpgrade,omitempty"`
	AutoMinorVersionUpgrade interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	BackupRetentionPeriod interface{} `yaml:"BackupRetentionPeriod,omitempty"`
	CharacterSetName interface{} `yaml:"CharacterSetName,omitempty"`
	CopyTagsToSnapshot interface{} `yaml:"CopyTagsToSnapshot,omitempty"`
	DBClusterIdentifier interface{} `yaml:"DBClusterIdentifier,omitempty"`
	DBInstanceClass interface{} `yaml:"DBInstanceClass"`
	DBInstanceIdentifier interface{} `yaml:"DBInstanceIdentifier,omitempty"`
	DBName interface{} `yaml:"DBName,omitempty"`
	DBParameterGroupName interface{} `yaml:"DBParameterGroupName,omitempty"`
	DBSnapshotIdentifier interface{} `yaml:"DBSnapshotIdentifier,omitempty"`
	DBSubnetGroupName interface{} `yaml:"DBSubnetGroupName,omitempty"`
	Domain interface{} `yaml:"Domain,omitempty"`
	DomainIAMRoleName interface{} `yaml:"DomainIAMRoleName,omitempty"`
	Engine interface{} `yaml:"Engine,omitempty"`
	EngineVersion interface{} `yaml:"EngineVersion,omitempty"`
	Iops interface{} `yaml:"Iops,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	LicenseModel interface{} `yaml:"LicenseModel,omitempty"`
	MasterUserPassword interface{} `yaml:"MasterUserPassword,omitempty"`
	MasterUsername interface{} `yaml:"MasterUsername,omitempty"`
	MonitoringInterval interface{} `yaml:"MonitoringInterval,omitempty"`
	MonitoringRoleArn interface{} `yaml:"MonitoringRoleArn,omitempty"`
	MultiAZ interface{} `yaml:"MultiAZ,omitempty"`
	OptionGroupName interface{} `yaml:"OptionGroupName,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	PreferredBackupWindow interface{} `yaml:"PreferredBackupWindow,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	PubliclyAccessible interface{} `yaml:"PubliclyAccessible,omitempty"`
	SourceDBInstanceIdentifier interface{} `yaml:"SourceDBInstanceIdentifier,omitempty"`
	SourceRegion interface{} `yaml:"SourceRegion,omitempty"`
	StorageEncrypted interface{} `yaml:"StorageEncrypted,omitempty"`
	StorageType interface{} `yaml:"StorageType,omitempty"`
	Timezone interface{} `yaml:"Timezone,omitempty"`
	DBSecurityGroups interface{} `yaml:"DBSecurityGroups,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VPCSecurityGroups interface{} `yaml:"VPCSecurityGroups,omitempty"`
}

func NewRDSDBInstance(properties RDSDBInstanceProperties, deps ...interface{}) RDSDBInstance {
	return RDSDBInstance{
		Type:       "AWS::RDS::DBInstance",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseRDSDBInstance(name string, data string) (cf types.ValueMap, err error) {
	var resource RDSDBInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: RDSDBInstance - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource RDSDBInstance) Validate() []error {
	return resource.Properties.Validate()
}

func (resource RDSDBInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.DBInstanceClass == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DBInstanceClass'"))
	}
	return errs
}
