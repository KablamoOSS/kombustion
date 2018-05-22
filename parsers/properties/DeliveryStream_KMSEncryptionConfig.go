package properties

	import "fmt"

type DeliveryStream_KMSEncryptionConfig struct {
	
	AWSKMSKeyARN interface{} `yaml:"AWSKMSKeyARN"`
}

func (resource DeliveryStream_KMSEncryptionConfig) Validate() []error {
	errs := []error{}
	
	if resource.AWSKMSKeyARN == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'AWSKMSKeyARN'"))
	}
	return errs
}
