package properties

	import "fmt"

type ScalingPolicy_TargetTrackingConfiguration struct {
	
	
	
	
	DisableScaleIn interface{} `yaml:"DisableScaleIn,omitempty"`
	TargetValue interface{} `yaml:"TargetValue"`
	PredefinedMetricSpecification *ScalingPolicy_PredefinedMetricSpecification `yaml:"PredefinedMetricSpecification,omitempty"`
	CustomizedMetricSpecification *ScalingPolicy_CustomizedMetricSpecification `yaml:"CustomizedMetricSpecification,omitempty"`
}

func (resource ScalingPolicy_TargetTrackingConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.TargetValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TargetValue'"))
	}
	return errs
}
