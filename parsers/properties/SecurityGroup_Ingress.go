package properties

	import "fmt"

type SecurityGroup_Ingress struct {
	
	
	
	
	
	
	
	
	
	CidrIp interface{} `yaml:"CidrIp,omitempty"`
	CidrIpv6 interface{} `yaml:"CidrIpv6,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	FromPort interface{} `yaml:"FromPort,omitempty"`
	IpProtocol interface{} `yaml:"IpProtocol"`
	SourceSecurityGroupId interface{} `yaml:"SourceSecurityGroupId,omitempty"`
	SourceSecurityGroupName interface{} `yaml:"SourceSecurityGroupName,omitempty"`
	SourceSecurityGroupOwnerId interface{} `yaml:"SourceSecurityGroupOwnerId,omitempty"`
	ToPort interface{} `yaml:"ToPort,omitempty"`
}

func (resource SecurityGroup_Ingress) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	if resource.IpProtocol == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'IpProtocol'"))
	}
	return errs
}
