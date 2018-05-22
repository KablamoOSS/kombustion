package properties

	import "fmt"

type User_LoginProfile struct {
	
	
	Password interface{} `yaml:"Password"`
	PasswordResetRequired interface{} `yaml:"PasswordResetRequired,omitempty"`
}

func (resource User_LoginProfile) Validate() []error {
	errs := []error{}
	
	
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	return errs
}
