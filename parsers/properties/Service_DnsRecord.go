package properties

	import "fmt"

type Service_DnsRecord struct {
	
	
	TTL interface{} `yaml:"TTL"`
	Type interface{} `yaml:"Type"`
}

func (resource Service_DnsRecord) Validate() []error {
	errs := []error{}
	
	
	if resource.TTL == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TTL'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
