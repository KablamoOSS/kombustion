package properties

	import "fmt"

type DeliveryStream_ExtendedS3DestinationConfiguration struct {
	
	
	
	
	
	
	
	
	
	
	BucketARN interface{} `yaml:"BucketARN"`
	CompressionFormat interface{} `yaml:"CompressionFormat"`
	Prefix interface{} `yaml:"Prefix"`
	RoleARN interface{} `yaml:"RoleARN"`
	S3BackupMode interface{} `yaml:"S3BackupMode,omitempty"`
	S3BackupConfiguration *DeliveryStream_S3DestinationConfiguration `yaml:"S3BackupConfiguration,omitempty"`
	ProcessingConfiguration *DeliveryStream_ProcessingConfiguration `yaml:"ProcessingConfiguration,omitempty"`
	EncryptionConfiguration *DeliveryStream_EncryptionConfiguration `yaml:"EncryptionConfiguration,omitempty"`
	CloudWatchLoggingOptions *DeliveryStream_CloudWatchLoggingOptions `yaml:"CloudWatchLoggingOptions,omitempty"`
	BufferingHints *DeliveryStream_BufferingHints `yaml:"BufferingHints"`
}

func (resource DeliveryStream_ExtendedS3DestinationConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	if resource.BucketARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketARN'"))
	}
	if resource.CompressionFormat == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CompressionFormat'"))
	}
	if resource.Prefix == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Prefix'"))
	}
	if resource.RoleARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'RoleARN'"))
	}
	if resource.BufferingHints == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BufferingHints'"))
	} else {
		errs = append(errs, resource.BufferingHints.Validate()...)
	}
	return errs
}
