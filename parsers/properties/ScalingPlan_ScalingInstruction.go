package properties

	import "fmt"

type ScalingPlan_ScalingInstruction struct {
	
	
	
	
	
	
	MaxCapacity interface{} `yaml:"MaxCapacity"`
	MinCapacity interface{} `yaml:"MinCapacity"`
	ResourceId interface{} `yaml:"ResourceId"`
	ScalableDimension interface{} `yaml:"ScalableDimension"`
	ServiceNamespace interface{} `yaml:"ServiceNamespace"`
	TargetTrackingConfigurations interface{} `yaml:"TargetTrackingConfigurations"`
}

func (resource ScalingPlan_ScalingInstruction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.MaxCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaxCapacity'"))
	}
	if resource.MinCapacity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinCapacity'"))
	}
	if resource.ResourceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ResourceId'"))
	}
	if resource.ScalableDimension == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ScalableDimension'"))
	}
	if resource.ServiceNamespace == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ServiceNamespace'"))
	}
	if resource.TargetTrackingConfigurations == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetTrackingConfigurations'"))
	}
	return errs
}
