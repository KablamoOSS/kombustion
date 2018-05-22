package properties

	import "fmt"

type Rule_Predicate struct {
	
	
	
	DataId interface{} `yaml:"DataId"`
	Negated interface{} `yaml:"Negated"`
	Type interface{} `yaml:"Type"`
}

func (resource Rule_Predicate) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DataId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DataId'"))
	}
	if resource.Negated == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Negated'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
