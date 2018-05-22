package properties


type DeploymentGroup_LoadBalancerInfo struct {
	
	
	ElbInfoList interface{} `yaml:"ElbInfoList,omitempty"`
	TargetGroupInfoList interface{} `yaml:"TargetGroupInfoList,omitempty"`
}

func (resource DeploymentGroup_LoadBalancerInfo) Validate() []error {
	errs := []error{}
	
	
	return errs
}
