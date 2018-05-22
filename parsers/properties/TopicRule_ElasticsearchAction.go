package properties

	import "fmt"

type TopicRule_ElasticsearchAction struct {
	
	
	
	
	
	Endpoint interface{} `yaml:"Endpoint"`
	Id interface{} `yaml:"Id"`
	Index interface{} `yaml:"Index"`
	RoleArn interface{} `yaml:"RoleArn"`
	Type interface{} `yaml:"Type"`
}

func (resource TopicRule_ElasticsearchAction) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Endpoint == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Endpoint'"))
	}
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	if resource.Index == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Index'"))
	}
	if resource.RoleArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleArn'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
