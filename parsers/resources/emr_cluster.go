package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EMRCluster struct {
	Type       string                      `yaml:"Type"`
	Properties EMRClusterProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EMRClusterProperties struct {
	AdditionalInfo interface{} `yaml:"AdditionalInfo,omitempty"`
	AutoScalingRole interface{} `yaml:"AutoScalingRole,omitempty"`
	CustomAmiId interface{} `yaml:"CustomAmiId,omitempty"`
	EbsRootVolumeSize interface{} `yaml:"EbsRootVolumeSize,omitempty"`
	JobFlowRole interface{} `yaml:"JobFlowRole"`
	LogUri interface{} `yaml:"LogUri,omitempty"`
	Name interface{} `yaml:"Name"`
	ReleaseLabel interface{} `yaml:"ReleaseLabel,omitempty"`
	ScaleDownBehavior interface{} `yaml:"ScaleDownBehavior,omitempty"`
	SecurityConfiguration interface{} `yaml:"SecurityConfiguration,omitempty"`
	ServiceRole interface{} `yaml:"ServiceRole"`
	VisibleToAllUsers interface{} `yaml:"VisibleToAllUsers,omitempty"`
	Applications interface{} `yaml:"Applications,omitempty"`
	BootstrapActions interface{} `yaml:"BootstrapActions,omitempty"`
	Configurations interface{} `yaml:"Configurations,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	Instances *properties.Cluster_JobFlowInstancesConfig `yaml:"Instances"`
}

func NewEMRCluster(properties EMRClusterProperties, deps ...interface{}) EMRCluster {
	return EMRCluster{
		Type:       "AWS::EMR::Cluster",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEMRCluster(name string, data string) (cf types.ValueMap, err error) {
	var resource EMRCluster
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EMRCluster - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EMRCluster) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EMRClusterProperties) Validate() []error {
	errs := []error{}
	if resource.JobFlowRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'JobFlowRole'"))
	}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.ServiceRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceRole'"))
	}
	if resource.Instances == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Instances'"))
	} else {
		errs = append(errs, resource.Instances.Validate()...)
	}
	return errs
}
