package properties

	import "fmt"

type Distribution_Cookies struct {
	
	
	Forward interface{} `yaml:"Forward"`
	WhitelistedNames interface{} `yaml:"WhitelistedNames,omitempty"`
}

func (resource Distribution_Cookies) Validate() []error {
	errs := []error{}
	
	
	if resource.Forward == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Forward'"))
	}
	return errs
}
