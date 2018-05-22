package properties

	import "fmt"

type Distribution_Origin struct {
	
	
	
	
	
	
	DomainName interface{} `yaml:"DomainName"`
	Id interface{} `yaml:"Id"`
	OriginPath interface{} `yaml:"OriginPath,omitempty"`
	S3OriginConfig *Distribution_S3OriginConfig `yaml:"S3OriginConfig,omitempty"`
	OriginCustomHeaders interface{} `yaml:"OriginCustomHeaders,omitempty"`
	CustomOriginConfig *Distribution_CustomOriginConfig `yaml:"CustomOriginConfig,omitempty"`
}

func (resource Distribution_Origin) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	if resource.Id == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Id'"))
	}
	return errs
}
