package properties

	import "fmt"

type RecordSetGroup_AliasTarget struct {
	
	
	
	DNSName interface{} `yaml:"DNSName"`
	EvaluateTargetHealth interface{} `yaml:"EvaluateTargetHealth,omitempty"`
	HostedZoneId interface{} `yaml:"HostedZoneId"`
}

func (resource RecordSetGroup_AliasTarget) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DNSName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DNSName'"))
	}
	if resource.HostedZoneId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HostedZoneId'"))
	}
	return errs
}
