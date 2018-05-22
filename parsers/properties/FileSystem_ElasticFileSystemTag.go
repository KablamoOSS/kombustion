package properties

	import "fmt"

type FileSystem_ElasticFileSystemTag struct {
	
	
	Key interface{} `yaml:"Key"`
	Value interface{} `yaml:"Value"`
}

func (resource FileSystem_ElasticFileSystemTag) Validate() []error {
	errs := []error{}
	
	
	if resource.Key == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Value'"))
	}
	return errs
}
