package properties

	import "fmt"

type SpotFleet_GroupIdentifier struct {
	
	GroupId interface{} `yaml:"GroupId"`
}

func (resource SpotFleet_GroupIdentifier) Validate() []error {
	errs := []error{}
	
	if resource.GroupId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GroupId'"))
	}
	return errs
}
