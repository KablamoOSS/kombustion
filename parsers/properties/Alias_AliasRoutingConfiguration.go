package properties

	import "fmt"

type Alias_AliasRoutingConfiguration struct {
	
	AdditionalVersionWeights interface{} `yaml:"AdditionalVersionWeights"`
}

func (resource Alias_AliasRoutingConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.AdditionalVersionWeights == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AdditionalVersionWeights'"))
	}
	return errs
}
