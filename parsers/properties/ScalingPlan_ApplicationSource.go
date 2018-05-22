package properties


type ScalingPlan_ApplicationSource struct {
	
	
	CloudFormationStackARN interface{} `yaml:"CloudFormationStackARN,omitempty"`
	TagFilters interface{} `yaml:"TagFilters,omitempty"`
}

func (resource ScalingPlan_ApplicationSource) Validate() []error {
	errs := []error{}
	
	
	return errs
}
