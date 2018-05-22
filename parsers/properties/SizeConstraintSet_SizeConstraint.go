package properties

	import "fmt"

type SizeConstraintSet_SizeConstraint struct {
	
	
	
	
	ComparisonOperator interface{} `yaml:"ComparisonOperator"`
	Size interface{} `yaml:"Size"`
	TextTransformation interface{} `yaml:"TextTransformation"`
	FieldToMatch *SizeConstraintSet_FieldToMatch `yaml:"FieldToMatch"`
}

func (resource SizeConstraintSet_SizeConstraint) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.ComparisonOperator == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ComparisonOperator'"))
	}
	if resource.Size == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Size'"))
	}
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
