package properties

	import "fmt"

type SimpleAD_VpcSettings struct {
	
	
	VpcId interface{} `yaml:"VpcId"`
	SubnetIds interface{} `yaml:"SubnetIds"`
}

func (resource SimpleAD_VpcSettings) Validate() []error {
	errs := []error{}
	
	
	if resource.VpcId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'VpcId'"))
	}
	if resource.SubnetIds == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'SubnetIds'"))
	}
	return errs
}
