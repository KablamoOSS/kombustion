package properties

	import "fmt"

type XssMatchSet_XssMatchTuple struct {
	
	
	TextTransformation interface{} `yaml:"TextTransformation"`
	FieldToMatch *XssMatchSet_FieldToMatch `yaml:"FieldToMatch"`
}

func (resource XssMatchSet_XssMatchTuple) Validate() []error {
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
