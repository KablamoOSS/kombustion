package properties

	import "fmt"

type ConfigurationTemplate_SourceConfiguration struct {
	
	
	ApplicationName interface{} `yaml:"ApplicationName"`
	TemplateName interface{} `yaml:"TemplateName"`
}

func (resource ConfigurationTemplate_SourceConfiguration) Validate() []error {
	errs := []error{}
	
	
	if resource.ApplicationName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ApplicationName'"))
	}
	if resource.TemplateName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TemplateName'"))
	}
	return errs
}
