package properties

	import "fmt"

type Service_DnsConfig struct {
	
	
	NamespaceId interface{} `yaml:"NamespaceId"`
	DnsRecords interface{} `yaml:"DnsRecords"`
}

func (resource Service_DnsConfig) Validate() []error {
	errs := []error{}
	
	
	if resource.NamespaceId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NamespaceId'"))
	}
	if resource.DnsRecords == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DnsRecords'"))
	}
	return errs
}
