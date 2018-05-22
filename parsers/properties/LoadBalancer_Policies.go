package properties

	import "fmt"

type LoadBalancer_Policies struct {
	
	
	
	
	
	PolicyName interface{} `yaml:"PolicyName"`
	PolicyType interface{} `yaml:"PolicyType"`
	Attributes interface{} `yaml:"Attributes"`
	InstancePorts interface{} `yaml:"InstancePorts,omitempty"`
	LoadBalancerPorts interface{} `yaml:"LoadBalancerPorts,omitempty"`
}

func (resource LoadBalancer_Policies) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.PolicyName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyName'"))
	}
	if resource.PolicyType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PolicyType'"))
	}
	if resource.Attributes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Attributes'"))
	}
	return errs
}
