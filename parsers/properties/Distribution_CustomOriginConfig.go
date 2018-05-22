package properties

	import "fmt"

type Distribution_CustomOriginConfig struct {
	
	
	
	
	
	
	HTTPPort interface{} `yaml:"HTTPPort,omitempty"`
	HTTPSPort interface{} `yaml:"HTTPSPort,omitempty"`
	OriginKeepaliveTimeout interface{} `yaml:"OriginKeepaliveTimeout,omitempty"`
	OriginProtocolPolicy interface{} `yaml:"OriginProtocolPolicy"`
	OriginReadTimeout interface{} `yaml:"OriginReadTimeout,omitempty"`
	OriginSSLProtocols interface{} `yaml:"OriginSSLProtocols,omitempty"`
}

func (resource Distribution_CustomOriginConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.OriginProtocolPolicy == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'OriginProtocolPolicy'"))
	}
	return errs
}
