package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DMSReplicationInstance struct {
	Type       string                      `yaml:"Type"`
	Properties DMSReplicationInstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSReplicationInstanceProperties struct {
	AllocatedStorage interface{} `yaml:"AllocatedStorage,omitempty"`
	AllowMajorVersionUpgrade interface{} `yaml:"AllowMajorVersionUpgrade,omitempty"`
	AutoMinorVersionUpgrade interface{} `yaml:"AutoMinorVersionUpgrade,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	EngineVersion interface{} `yaml:"EngineVersion,omitempty"`
	KmsKeyId interface{} `yaml:"KmsKeyId,omitempty"`
	MultiAZ interface{} `yaml:"MultiAZ,omitempty"`
	PreferredMaintenanceWindow interface{} `yaml:"PreferredMaintenanceWindow,omitempty"`
	PubliclyAccessible interface{} `yaml:"PubliclyAccessible,omitempty"`
	ReplicationInstanceClass interface{} `yaml:"ReplicationInstanceClass"`
	ReplicationInstanceIdentifier interface{} `yaml:"ReplicationInstanceIdentifier,omitempty"`
	ReplicationSubnetGroupIdentifier interface{} `yaml:"ReplicationSubnetGroupIdentifier,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	VpcSecurityGroupIds interface{} `yaml:"VpcSecurityGroupIds,omitempty"`
}

func NewDMSReplicationInstance(properties DMSReplicationInstanceProperties, deps ...interface{}) DMSReplicationInstance {
	return DMSReplicationInstance{
		Type:       "AWS::DMS::ReplicationInstance",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSReplicationInstance(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSReplicationInstance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSReplicationInstance - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSReplicationInstance) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSReplicationInstanceProperties) Validate() []error {
	errs := []error{}
	if resource.ReplicationInstanceClass == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationInstanceClass'"))
	}
	return errs
}
