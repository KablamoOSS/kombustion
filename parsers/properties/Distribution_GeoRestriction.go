package properties

	import "fmt"

type Distribution_GeoRestriction struct {
	
	
	RestrictionType interface{} `yaml:"RestrictionType"`
	Locations interface{} `yaml:"Locations,omitempty"`
}

func (resource Distribution_GeoRestriction) Validate() []error {
	errs := []error{}
	
	
	if resource.RestrictionType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RestrictionType'"))
	}
	return errs
}
