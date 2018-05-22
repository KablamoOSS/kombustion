package properties

	import "fmt"

type Stream_StreamEncryption struct {
	
	
	EncryptionType interface{} `yaml:"EncryptionType"`
	KeyId interface{} `yaml:"KeyId"`
}

func (resource Stream_StreamEncryption) Validate() []error {
	errs := []error{}
	
	
	if resource.EncryptionType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'EncryptionType'"))
	}
	if resource.KeyId == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'KeyId'"))
	}
	return errs
}
