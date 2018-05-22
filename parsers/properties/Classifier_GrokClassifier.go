package properties

	import "fmt"

type Classifier_GrokClassifier struct {
	
	
	
	
	Classification interface{} `yaml:"Classification"`
	CustomPatterns interface{} `yaml:"CustomPatterns,omitempty"`
	GrokPattern interface{} `yaml:"GrokPattern"`
	Name interface{} `yaml:"Name,omitempty"`
}

func (resource Classifier_GrokClassifier) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.Classification == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Classification'"))
	}
	if resource.GrokPattern == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'GrokPattern'"))
	}
	return errs
}
