package properties

	import "fmt"

type Stack_RdsDbInstance struct {
	
	
	
	DbPassword interface{} `yaml:"DbPassword"`
	DbUser interface{} `yaml:"DbUser"`
	RdsDbInstanceArn interface{} `yaml:"RdsDbInstanceArn"`
}

func (resource Stack_RdsDbInstance) Validate() []error {
	errs := []error{}
	
	
	
	if resource.DbPassword == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DbPassword'"))
	}
	if resource.DbUser == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DbUser'"))
	}
	if resource.RdsDbInstanceArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RdsDbInstanceArn'"))
	}
	return errs
}
