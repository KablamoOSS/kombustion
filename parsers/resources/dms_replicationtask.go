package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
)

type DMSReplicationTask struct {
	Type       string                      `yaml:"Type"`
	Properties DMSReplicationTaskProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type DMSReplicationTaskProperties struct {
	CdcStartTime interface{} `yaml:"CdcStartTime,omitempty"`
	MigrationType interface{} `yaml:"MigrationType"`
	ReplicationInstanceArn interface{} `yaml:"ReplicationInstanceArn"`
	ReplicationTaskIdentifier interface{} `yaml:"ReplicationTaskIdentifier,omitempty"`
	ReplicationTaskSettings interface{} `yaml:"ReplicationTaskSettings,omitempty"`
	SourceEndpointArn interface{} `yaml:"SourceEndpointArn"`
	TableMappings interface{} `yaml:"TableMappings"`
	TargetEndpointArn interface{} `yaml:"TargetEndpointArn"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewDMSReplicationTask(properties DMSReplicationTaskProperties, deps ...interface{}) DMSReplicationTask {
	return DMSReplicationTask{
		Type:       "AWS::DMS::ReplicationTask",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseDMSReplicationTask(name string, data string) (cf types.ValueMap, err error) {
	var resource DMSReplicationTask
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: DMSReplicationTask - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource DMSReplicationTask) Validate() []error {
	return resource.Properties.Validate()
}

func (resource DMSReplicationTaskProperties) Validate() []error {
	errs := []error{}
	if resource.MigrationType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MigrationType'"))
	}
	if resource.ReplicationInstanceArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicationInstanceArn'"))
	}
	if resource.SourceEndpointArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SourceEndpointArn'"))
	}
	if resource.TableMappings == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TableMappings'"))
	}
	if resource.TargetEndpointArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetEndpointArn'"))
	}
	return errs
}
