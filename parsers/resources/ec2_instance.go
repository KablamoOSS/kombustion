package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type EC2Instance struct {
	Type       string                      `yaml:"Type"`
	Properties EC2InstanceProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type EC2InstanceProperties struct {
	AdditionalInfo interface{} `yaml:"AdditionalInfo,omitempty"`
	Affinity interface{} `yaml:"Affinity,omitempty"`
	AvailabilityZone interface{} `yaml:"AvailabilityZone,omitempty"`
	DisableApiTermination interface{} `yaml:"DisableApiTermination,omitempty"`
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	HostId interface{} `yaml:"HostId,omitempty"`
	IamInstanceProfile interface{} `yaml:"IamInstanceProfile,omitempty"`
	ImageId interface{} `yaml:"ImageId"`
	InstanceInitiatedShutdownBehavior interface{} `yaml:"InstanceInitiatedShutdownBehavior,omitempty"`
	InstanceType interface{} `yaml:"InstanceType,omitempty"`
	Ipv6AddressCount interface{} `yaml:"Ipv6AddressCount,omitempty"`
	KernelId interface{} `yaml:"KernelId,omitempty"`
	KeyName interface{} `yaml:"KeyName,omitempty"`
	Monitoring interface{} `yaml:"Monitoring,omitempty"`
	PlacementGroupName interface{} `yaml:"PlacementGroupName,omitempty"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress,omitempty"`
	RamdiskId interface{} `yaml:"RamdiskId,omitempty"`
	SourceDestCheck interface{} `yaml:"SourceDestCheck,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	Tenancy interface{} `yaml:"Tenancy,omitempty"`
	UserData interface{} `yaml:"UserData,omitempty"`
	BlockDeviceMappings interface{} `yaml:"BlockDeviceMappings,omitempty"`
	SsmAssociations interface{} `yaml:"SsmAssociations,omitempty"`
	ElasticGpuSpecifications interface{} `yaml:"ElasticGpuSpecifications,omitempty"`
	Ipv6Addresses interface{} `yaml:"Ipv6Addresses,omitempty"`
	NetworkInterfaces interface{} `yaml:"NetworkInterfaces,omitempty"`
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	Volumes interface{} `yaml:"Volumes,omitempty"`
	CreditSpecification *properties.Instance_CreditSpecification `yaml:"CreditSpecification,omitempty"`
}

func NewEC2Instance(properties EC2InstanceProperties, deps ...interface{}) EC2Instance {
	return EC2Instance{
		Type:       "AWS::EC2::Instance",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseEC2Instance(name string, data string) (cf types.ValueMap, err error) {
	var resource EC2Instance
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: EC2Instance - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource EC2Instance) Validate() []error {
	return resource.Properties.Validate()
}

func (resource EC2InstanceProperties) Validate() []error {
	errs := []error{}
	if resource.ImageId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ImageId'"))
	}
	return errs
}
