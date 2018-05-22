package properties

	import "fmt"

type ByteMatchSet_ByteMatchTuple struct {
	
	
	
	
	
	PositionalConstraint interface{} `yaml:"PositionalConstraint"`
	TargetString interface{} `yaml:"TargetString,omitempty"`
	TargetStringBase64 interface{} `yaml:"TargetStringBase64,omitempty"`
	TextTransformation interface{} `yaml:"TextTransformation"`
	FieldToMatch *ByteMatchSet_FieldToMatch `yaml:"FieldToMatch"`
}

func (resource ByteMatchSet_ByteMatchTuple) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.PositionalConstraint == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'PositionalConstraint'"))
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
