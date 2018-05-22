package properties

	import "fmt"

type MaintenanceWindowTask_Target struct {
	
	
	Key interface{} `yaml:"Key"`
	Values interface{} `yaml:"Values,omitempty"`
}

func (resource MaintenanceWindowTask_Target) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	return errs
}
