package properties

	import "fmt"

type SpotFleet_SpotFleetLaunchSpecification struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	EbsOptimized interface{} `yaml:"EbsOptimized,omitempty"`
	ImageId interface{} `yaml:"ImageId"`
	InstanceType interface{} `yaml:"InstanceType"`
	KernelId interface{} `yaml:"KernelId,omitempty"`
	KeyName interface{} `yaml:"KeyName,omitempty"`
	RamdiskId interface{} `yaml:"RamdiskId,omitempty"`
	SpotPrice interface{} `yaml:"SpotPrice,omitempty"`
	SubnetId interface{} `yaml:"SubnetId,omitempty"`
	UserData interface{} `yaml:"UserData,omitempty"`
	WeightedCapacity interface{} `yaml:"WeightedCapacity,omitempty"`
	Placement *SpotFleet_SpotPlacement `yaml:"Placement,omitempty"`
	Monitoring *SpotFleet_SpotFleetMonitoring `yaml:"Monitoring,omitempty"`
	BlockDeviceMappings interface{} `yaml:"BlockDeviceMappings,omitempty"`
	NetworkInterfaces interface{} `yaml:"NetworkInterfaces,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	TagSpecifications interface{} `yaml:"TagSpecifications,omitempty"`
	IamInstanceProfile *SpotFleet_IamInstanceProfileSpecification `yaml:"IamInstanceProfile,omitempty"`
}

func (resource SpotFleet_SpotFleetLaunchSpecification) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.ImageId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ImageId'"))
	}
	if resource.InstanceType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstanceType'"))
	}
	return errs
}
