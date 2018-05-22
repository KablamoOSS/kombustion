package properties

	import "fmt"

type HostedZone_HostedZoneTag struct {
	
	
	Key interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

func (resource HostedZone_HostedZoneTag) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
