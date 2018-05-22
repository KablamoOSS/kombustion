package properties

	import "fmt"

type Bucket_Transition struct {
	
	
	
	StorageClass interface{} `yaml:"StorageClass"`
	TransitionDate interface{} `yaml:"TransitionDate,omitempty"`
	TransitionInDays interface{} `yaml:"TransitionInDays,omitempty"`
}

func (resource Bucket_Transition) Validate() []error {
	errs := []error{}
	
	
	
	if resource.StorageClass == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'StorageClass'"))
	}
	return errs
}
