package properties

	import "fmt"

type Distribution_Restrictions struct {
	
	GeoRestriction *Distribution_GeoRestriction `yaml:"GeoRestriction"`
}

func (resource Distribution_Restrictions) Validate() []error {
	errs := []error{}
	
	if resource.GeoRestriction == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GeoRestriction'"))
	} else {
		errs = append(errs, resource.GeoRestriction.Validate()...)
	}
	return errs
}
