package properties


type Trail_EventSelector struct {
	
	
	
	IncludeManagementEvents interface{} `yaml:"IncludeManagementEvents,omitempty"`
	ReadWriteType interface{} `yaml:"ReadWriteType,omitempty"`
	DataResources interface{} `yaml:"DataResources,omitempty"`
}

func (resource Trail_EventSelector) Validate() []error {
	errs := []error{}
	
	
	
	return errs
}
