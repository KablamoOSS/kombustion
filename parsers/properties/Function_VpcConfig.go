package properties

	import "fmt"

type Function_VpcConfig struct {
	
	
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

func (resource Function_VpcConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.SecurityGroupIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SecurityGroupIds'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
