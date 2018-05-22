package properties

	import "fmt"

type ReceiptFilter_IpFilter struct {
	
	
	Cidr interface{} `yaml:"Cidr"`
	Policy interface{} `yaml:"Policy"`
}

func (resource ReceiptFilter_IpFilter) Validate() []error {
	errs := []error{}
	
	
	if resource.Cidr == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Cidr'"))
	}
	if resource.Policy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Policy'"))
	}
	return errs
}
