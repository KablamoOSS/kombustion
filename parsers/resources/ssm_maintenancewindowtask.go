package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type SSMMaintenanceWindowTask struct {
	Type       string                      `yaml:"Type"`
	Properties SSMMaintenanceWindowTaskProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SSMMaintenanceWindowTaskProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	MaxConcurrency interface{} `yaml:"MaxConcurrency"`
	MaxErrors interface{} `yaml:"MaxErrors"`
	Name interface{} `yaml:"Name,omitempty"`
	Priority interface{} `yaml:"Priority"`
	ServiceRoleArn interface{} `yaml:"ServiceRoleArn"`
	TaskArn interface{} `yaml:"TaskArn"`
	TaskParameters interface{} `yaml:"TaskParameters,omitempty"`
	TaskType interface{} `yaml:"TaskType"`
	WindowId interface{} `yaml:"WindowId,omitempty"`
	TaskInvocationParameters *properties.MaintenanceWindowTask_TaskInvocationParameters `yaml:"TaskInvocationParameters,omitempty"`
	LoggingInfo *properties.MaintenanceWindowTask_LoggingInfo `yaml:"LoggingInfo,omitempty"`
	Targets interface{} `yaml:"Targets"`
}

func NewSSMMaintenanceWindowTask(properties SSMMaintenanceWindowTaskProperties, deps ...interface{}) SSMMaintenanceWindowTask {
	return SSMMaintenanceWindowTask{
		Type:       "AWS::SSM::MaintenanceWindowTask",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMMaintenanceWindowTask(name string, data string) (cf types.ValueMap, err error) {
	var resource SSMMaintenanceWindowTask
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMMaintenanceWindowTask - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SSMMaintenanceWindowTask) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMMaintenanceWindowTaskProperties) Validate() []error {
	errs := []error{}
	if resource.MaxConcurrency == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxConcurrency'"))
	}
	if resource.MaxErrors == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxErrors'"))
	}
	if resource.Priority == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Priority'"))
	}
	if resource.ServiceRoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRoleArn'"))
	}
	if resource.TaskArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TaskArn'"))
	}
	if resource.TaskType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TaskType'"))
	}
	if resource.Targets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Targets'"))
	}
	return errs
}
