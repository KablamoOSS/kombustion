package properties

	import "fmt"

type OptionGroup_OptionConfiguration struct {
	
	
	
	
	
	
	OptionName interface{} `yaml:"OptionName"`
	OptionVersion interface{} `yaml:"OptionVersion,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	OptionSettings *OptionGroup_OptionSetting `yaml:"OptionSettings,omitempty"`
	DBSecurityGroupMemberships interface{} `yaml:"DBSecurityGroupMemberships,omitempty"`
	VpcSecurityGroupMemberships interface{} `yaml:"VpcSecurityGroupMemberships,omitempty"`
}

func (resource OptionGroup_OptionConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.OptionName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OptionName'"))
	}
	return errs
}
