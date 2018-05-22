package properties


type ReceiptRule_Rule struct {
	
	
	
	
	
	
	Enabled interface{} `yaml:"Enabled,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	ScanEnabled interface{} `yaml:"ScanEnabled,omitempty"`
	TlsPolicy interface{} `yaml:"TlsPolicy,omitempty"`
	Actions interface{} `yaml:"Actions,omitempty"`
	Recipients interface{} `yaml:"Recipients,omitempty"`
}

func (resource ReceiptRule_Rule) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	return errs
}
