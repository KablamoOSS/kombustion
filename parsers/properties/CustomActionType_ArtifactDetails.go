package properties

	import "fmt"

type CustomActionType_ArtifactDetails struct {
	
	
	MaximumCount interface{} `yaml:"MaximumCount"`
	MinimumCount interface{} `yaml:"MinimumCount"`
}

func (resource CustomActionType_ArtifactDetails) Validate() []error {
	errs := []error{}
	
	
	if resource.MaximumCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MaximumCount'"))
	}
	if resource.MinimumCount == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MinimumCount'"))
	}
	return errs
}
