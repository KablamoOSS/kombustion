package properties

	import "fmt"

type LoadBalancer_Listeners struct {
	
	
	
	
	
	
	InstancePort interface{} `yaml:"InstancePort"`
	InstanceProtocol interface{} `yaml:"InstanceProtocol,omitempty"`
	LoadBalancerPort interface{} `yaml:"LoadBalancerPort"`
	Protocol interface{} `yaml:"Protocol"`
	SSLCertificateId interface{} `yaml:"SSLCertificateId,omitempty"`
	PolicyNames interface{} `yaml:"PolicyNames,omitempty"`
}

func (resource LoadBalancer_Listeners) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.InstancePort == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'InstancePort'"))
	}
	if resource.LoadBalancerPort == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'LoadBalancerPort'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	return errs
}
