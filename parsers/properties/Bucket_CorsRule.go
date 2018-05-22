package properties

	import "fmt"

type Bucket_CorsRule struct {
	
	
	
	
	
	
	Id interface{} `yaml:"Id,omitempty"`
	MaxAge interface{} `yaml:"MaxAge,omitempty"`
	AllowedHeaders interface{} `yaml:"AllowedHeaders,omitempty"`
	AllowedMethods interface{} `yaml:"AllowedMethods"`
	AllowedOrigins interface{} `yaml:"AllowedOrigins"`
	ExposedHeaders interface{} `yaml:"ExposedHeaders,omitempty"`
}

func (resource Bucket_CorsRule) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.AllowedMethods == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AllowedMethods'"))
	}
	if resource.AllowedOrigins == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AllowedOrigins'"))
	}
	return errs
}
