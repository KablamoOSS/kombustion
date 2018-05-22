package properties

	import "fmt"

type Pipeline_ArtifactStore struct {
	
	
	
	Location interface{} `yaml:"Location"`
	Type interface{} `yaml:"Type"`
	EncryptionKey *Pipeline_EncryptionKey `yaml:"EncryptionKey,omitempty"`
}

func (resource Pipeline_ArtifactStore) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Location == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Location'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
