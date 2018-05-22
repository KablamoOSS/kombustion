package properties

	import "fmt"

type Service_LoadBalancer struct {
	
	
	
	
	ContainerName interface{} `yaml:"ContainerName,omitempty"`
	ContainerPort interface{} `yaml:"ContainerPort"`
	LoadBalancerName interface{} `yaml:"LoadBalancerName,omitempty"`
	TargetGroupArn interface{} `yaml:"TargetGroupArn,omitempty"`
}

func (resource Service_LoadBalancer) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.ContainerPort == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ContainerPort'"))
	}
	return errs
}
