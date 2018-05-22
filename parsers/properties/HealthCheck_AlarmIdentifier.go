package properties

	import "fmt"

type HealthCheck_AlarmIdentifier struct {
	
	
	Name interface{} `yaml:"Name"`
	Region interface{} `yaml:"Region"`
}

func (resource HealthCheck_AlarmIdentifier) Validate() []error {
	errs := []error{}
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Region == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Region'"))
	}
	return errs
}
