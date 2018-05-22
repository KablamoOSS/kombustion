package properties

	import "fmt"

type ClusterParameterGroup_Parameter struct {
	
	
	ParameterName interface{} `yaml:"ParameterName"`
	ParameterValue interface{} `yaml:"ParameterValue"`
}

func (resource ClusterParameterGroup_Parameter) Validate() []error {
	errs := []error{}
	
	
	if resource.ParameterName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterName'"))
	}
	if resource.ParameterValue == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ParameterValue'"))
	}
	return errs
}
