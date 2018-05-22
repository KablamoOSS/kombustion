package properties

	import "fmt"

type TargetGroup_Matcher struct {
	
	HttpCode interface{} `yaml:"HttpCode"`
}

func (resource TargetGroup_Matcher) Validate() []error {
	errs := []error{}
	
	if resource.HttpCode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'HttpCode'"))
	}
	return errs
}
