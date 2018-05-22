package properties

	import "fmt"

type Fleet_IpPermission struct {
	
	
	
	
	FromPort interface{} `yaml:"FromPort"`
	IpRange interface{} `yaml:"IpRange"`
	Protocol interface{} `yaml:"Protocol"`
	ToPort interface{} `yaml:"ToPort"`
}

func (resource Fleet_IpPermission) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.FromPort == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FromPort'"))
	}
	if resource.IpRange == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpRange'"))
	}
	if resource.Protocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Protocol'"))
	}
	if resource.ToPort == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ToPort'"))
	}
	return errs
}
