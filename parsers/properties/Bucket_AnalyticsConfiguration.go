package properties

	import "fmt"

type Bucket_AnalyticsConfiguration struct {
	
	
	
	
	Id interface{} `yaml:"Id"`
	Prefix interface{} `yaml:"Prefix,omitempty"`
	StorageClassAnalysis *Bucket_StorageClassAnalysis `yaml:"StorageClassAnalysis"`
	TagFilters interface{} `yaml:"TagFilters,omitempty"`
}

func (resource Bucket_AnalyticsConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.StorageClassAnalysis == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StorageClassAnalysis'"))
	} else {
		errs = append(errs, resource.StorageClassAnalysis.Validate()...)
	}
	return errs
}
