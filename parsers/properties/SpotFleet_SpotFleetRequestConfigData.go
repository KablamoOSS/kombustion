package properties

	import "fmt"

type SpotFleet_SpotFleetRequestConfigData struct {
	
	
	
	
	
	
	
	
	
	
	
	
	AllocationStrategy interface{} `yaml:"AllocationStrategy,omitempty"`
	ExcessCapacityTerminationPolicy interface{} `yaml:"ExcessCapacityTerminationPolicy,omitempty"`
	IamFleetRole interface{} `yaml:"IamFleetRole"`
	ReplaceUnhealthyInstances interface{} `yaml:"ReplaceUnhealthyInstances,omitempty"`
	SpotPrice interface{} `yaml:"SpotPrice,omitempty"`
	TargetCapacity interface{} `yaml:"TargetCapacity"`
	TerminateInstancesWithExpiration interface{} `yaml:"TerminateInstancesWithExpiration,omitempty"`
	Type interface{} `yaml:"Type,omitempty"`
	ValidFrom interface{} `yaml:"ValidFrom,omitempty"`
	ValidUntil interface{} `yaml:"ValidUntil,omitempty"`
	LaunchSpecifications interface{} `yaml:"LaunchSpecifications,omitempty"`
	LaunchTemplateConfigs interface{} `yaml:"LaunchTemplateConfigs,omitempty"`
}

func (resource SpotFleet_SpotFleetRequestConfigData) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.IamFleetRole == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IamFleetRole'"))
	}
	if resource.TargetCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetCapacity'"))
	}
	return errs
}
