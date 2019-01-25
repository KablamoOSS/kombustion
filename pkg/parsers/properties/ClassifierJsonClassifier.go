package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ClassifierJsonClassifier Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-classifier-jsonclassifier.html
type ClassifierJsonClassifier struct {
	JsonPath interface{} `yaml:"JsonPath"`
	Name     interface{} `yaml:"Name,omitempty"`
}

// ClassifierJsonClassifier validation
func (resource ClassifierJsonClassifier) Validate() []error {
	errors := []error{}

	return errors
}