package properties

	import "fmt"

type DeliveryStream_RedshiftDestinationConfiguration struct {
	
	
	
	
	
	
	
	
	ClusterJDBCURL interface{} `yaml:"ClusterJDBCURL"`
	Password interface{} `yaml:"Password"`
	RoleARN interface{} `yaml:"RoleARN"`
	Username interface{} `yaml:"Username"`
	S3Configuration *DeliveryStream_S3DestinationConfiguration `yaml:"S3Configuration"`
	ProcessingConfiguration *DeliveryStream_ProcessingConfiguration `yaml:"ProcessingConfiguration,omitempty"`
	CopyCommand *DeliveryStream_CopyCommand `yaml:"CopyCommand"`
	CloudWatchLoggingOptions *DeliveryStream_CloudWatchLoggingOptions `yaml:"CloudWatchLoggingOptions,omitempty"`
}

func (resource DeliveryStream_RedshiftDestinationConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	if resource.ClusterJDBCURL == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ClusterJDBCURL'"))
	}
	if resource.Password == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Password'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.Username == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Username'"))
	}
	if resource.S3Configuration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Configuration'"))
	} else {
		errs = append(errs, resource.S3Configuration.Validate()...)
	}
	if resource.CopyCommand == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CopyCommand'"))
	} else {
		errs = append(errs, resource.CopyCommand.Validate()...)
	}
	return errs
}
