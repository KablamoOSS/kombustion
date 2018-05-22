package properties

	import "fmt"

type Bucket_ReplicationDestination struct {
	
	
	
	
	
	Account interface{} `yaml:"Account,omitempty"`
	Bucket interface{} `yaml:"Bucket"`
	StorageClass interface{} `yaml:"StorageClass,omitempty"`
	EncryptionConfiguration *Bucket_EncryptionConfiguration `yaml:"EncryptionConfiguration,omitempty"`
	AccessControlTranslation *Bucket_AccessControlTranslation `yaml:"AccessControlTranslation,omitempty"`
}

func (resource Bucket_ReplicationDestination) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.Bucket == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Bucket'"))
	}
	return errs
}
