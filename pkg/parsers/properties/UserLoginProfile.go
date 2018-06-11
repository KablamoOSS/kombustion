package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// UserLoginProfile Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
type UserLoginProfile struct {
	Password              interface{} `yaml:"Password"`
	PasswordResetRequired interface{} `yaml:"PasswordResetRequired,omitempty"`
}

func (resource UserLoginProfile) Validate() []error {
	errs := []error{}

	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	return errs
}
