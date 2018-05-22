package properties


type Endpoint_MongoDbSettings struct {
	
	
	
	
	
	
	
	
	
	
	
	AuthMechanism interface{} `yaml:"AuthMechanism,omitempty"`
	AuthSource interface{} `yaml:"AuthSource,omitempty"`
	AuthType interface{} `yaml:"AuthType,omitempty"`
	DatabaseName interface{} `yaml:"DatabaseName,omitempty"`
	DocsToInvestigate interface{} `yaml:"DocsToInvestigate,omitempty"`
	ExtractDocId interface{} `yaml:"ExtractDocId,omitempty"`
	NestingLevel interface{} `yaml:"NestingLevel,omitempty"`
	Password interface{} `yaml:"Password,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	ServerName interface{} `yaml:"ServerName,omitempty"`
	Username interface{} `yaml:"Username,omitempty"`
}

func (resource Endpoint_MongoDbSettings) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
