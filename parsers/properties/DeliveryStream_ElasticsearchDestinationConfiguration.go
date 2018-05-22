package properties

	import "fmt"

type DeliveryStream_ElasticsearchDestinationConfiguration struct {
	
	
	
	
	
	
	
	
	
	
	
	DomainARN interface{} `yaml:"DomainARN"`
	IndexName interface{} `yaml:"IndexName"`
	IndexRotationPeriod interface{} `yaml:"IndexRotationPeriod"`
	RoleARN interface{} `yaml:"RoleARN"`
	S3BackupMode interface{} `yaml:"S3BackupMode"`
	TypeName interface{} `yaml:"TypeName"`
	S3Configuration *DeliveryStream_S3DestinationConfiguration `yaml:"S3Configuration"`
	ProcessingConfiguration *DeliveryStream_ProcessingConfiguration `yaml:"ProcessingConfiguration,omitempty"`
	RetryOptions *DeliveryStream_ElasticsearchRetryOptions `yaml:"RetryOptions"`
	BufferingHints *DeliveryStream_ElasticsearchBufferingHints `yaml:"BufferingHints"`
	CloudWatchLoggingOptions *DeliveryStream_CloudWatchLoggingOptions `yaml:"CloudWatchLoggingOptions,omitempty"`
}

func (resource DeliveryStream_ElasticsearchDestinationConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	if resource.DomainARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainARN'"))
	}
	if resource.IndexName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IndexName'"))
	}
	if resource.IndexRotationPeriod == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IndexRotationPeriod'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.S3BackupMode == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3BackupMode'"))
	}
	if resource.TypeName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'TypeName'"))
	}
	if resource.S3Configuration == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3Configuration'"))
	} else {
		errs = append(errs, resource.S3Configuration.Validate()...)
	}
	if resource.RetryOptions == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RetryOptions'"))
	} else {
		errs = append(errs, resource.RetryOptions.Validate()...)
	}
	if resource.BufferingHints == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BufferingHints'"))
	} else {
		errs = append(errs, resource.BufferingHints.Validate()...)
	}
	return errs
}
