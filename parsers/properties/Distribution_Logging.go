package properties

	import "fmt"

type Distribution_Logging struct {
	
	
	
	Bucket interface{} `yaml:"Bucket"`
	IncludeCookies interface{} `yaml:"IncludeCookies,omitempty"`
	Prefix interface{} `yaml:"Prefix,omitempty"`
}

func (resource Distribution_Logging) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	return errs
}
