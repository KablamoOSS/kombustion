package properties

	import "fmt"

type StreamingDistribution_S3Origin struct {
	
	
	DomainName interface{} `yaml:"DomainName"`
	OriginAccessIdentity interface{} `yaml:"OriginAccessIdentity"`
}

func (resource StreamingDistribution_S3Origin) Validate() []error {
	errs := []error{}
	
	
	if resource.DomainName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'DomainName'"))
	}
	if resource.OriginAccessIdentity == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OriginAccessIdentity'"))
	}
	return errs
}
