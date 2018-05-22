package properties

	import "fmt"

type ReceiptFilter_Filter struct {
	
	
	Name interface{} `yaml:"Name,omitempty"`
	IpFilter *ReceiptFilter_IpFilter `yaml:"IpFilter"`
}

func (resource ReceiptFilter_Filter) Validate() []error {
	errs := []error{}
	
	
	if resource.IpFilter == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpFilter'"))
	} else {
		errs = append(errs, resource.IpFilter.Validate()...)
	}
	return errs
}
