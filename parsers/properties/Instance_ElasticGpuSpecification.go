package properties

	import "fmt"

type Instance_ElasticGpuSpecification struct {
	
	Type interface{} `yaml:"Type"`
}

func (resource Instance_ElasticGpuSpecification) Validate() []error {
	errs := []error{}
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
