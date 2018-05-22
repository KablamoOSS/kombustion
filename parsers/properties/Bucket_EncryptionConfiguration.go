package properties

	import "fmt"

type Bucket_EncryptionConfiguration struct {
	
	ReplicaKmsKeyID interface{} `yaml:"ReplicaKmsKeyID"`
}

func (resource Bucket_EncryptionConfiguration) Validate() []error {
	errs := []error{}
	
	if resource.ReplicaKmsKeyID == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ReplicaKmsKeyID'"))
	}
	return errs
}
