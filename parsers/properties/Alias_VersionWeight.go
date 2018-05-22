package properties

	import "fmt"

type Alias_VersionWeight struct {
	
	
	FunctionVersion interface{} `yaml:"FunctionVersion"`
	FunctionWeight interface{} `yaml:"FunctionWeight"`
}

func (resource Alias_VersionWeight) Validate() []error {
	errs := []error{}
	
	
	if resource.FunctionVersion == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionVersion'"))
	}
	if resource.FunctionWeight == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FunctionWeight'"))
	}
	return errs
}
