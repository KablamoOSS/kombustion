package properties

	import "fmt"

type Bucket_CorsConfiguration struct {
	
	CorsRules interface{} `yaml:"CorsRules"`
}

func (resource Bucket_CorsConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.CorsRules == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CorsRules'"))
	}
	return errs
}
