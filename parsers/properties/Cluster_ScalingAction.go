package properties

	import "fmt"

type Cluster_ScalingAction struct {
	
	
	Market interface{} `yaml:"Market,omitempty"`
	SimpleScalingPolicyConfiguration *Cluster_SimpleScalingPolicyConfiguration `yaml:"SimpleScalingPolicyConfiguration"`
}

func (resource Cluster_ScalingAction) Validate() []error {
	errs := []error{}
	
	
	if resource.SimpleScalingPolicyConfiguration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SimpleScalingPolicyConfiguration'"))
	} else {
		errs = append(errs, resource.SimpleScalingPolicyConfiguration.Validate()...)
	}
	return errs
}
