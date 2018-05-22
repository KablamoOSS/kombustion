package properties

	import "fmt"

type DeliveryStream_S3DestinationConfiguration struct {
	
	
	
	
	
	
	
	BucketARN interface{} `yaml:"BucketARN"`
	CompressionFormat interface{} `yaml:"CompressionFormat"`
	Prefix interface{} `yaml:"Prefix,omitempty"`
	RoleARN interface{} `yaml:"RoleARN"`
	EncryptionConfiguration *DeliveryStream_EncryptionConfiguration `yaml:"EncryptionConfiguration,omitempty"`
	CloudWatchLoggingOptions *DeliveryStream_CloudWatchLoggingOptions `yaml:"CloudWatchLoggingOptions,omitempty"`
	BufferingHints *DeliveryStream_BufferingHints `yaml:"BufferingHints"`
}

func (resource DeliveryStream_S3DestinationConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	if resource.BucketARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketARN'"))
	}
	if resource.CompressionFormat == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'CompressionFormat'"))
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
