package properties

	import "fmt"

type SqlInjectionMatchSet_SqlInjectionMatchTuple struct {
	
	
	TextTransformation interface{} `yaml:"TextTransformation"`
	FieldToMatch *SqlInjectionMatchSet_FieldToMatch `yaml:"FieldToMatch"`
}

func (resource SqlInjectionMatchSet_SqlInjectionMatchTuple) Validate() []error {
	errs := []error{}
	
	
	if resource.TextTransformation == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TextTransformation'"))
	}
	if resource.FieldToMatch == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'FieldToMatch'"))
	} else {
		errs = append(errs, resource.FieldToMatch.Validate()...)
	}
	return errs
}
