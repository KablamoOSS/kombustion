package properties

	import "fmt"

type Service_AwsVpcConfiguration struct {
	
	
	
	AssignPublicIp interface{} `yaml:"AssignPublicIp,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	Subnets interface{} `yaml:"Subnets"`
}

func (resource Service_AwsVpcConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Subnets == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Subnets'"))
	}
	return errs
}
