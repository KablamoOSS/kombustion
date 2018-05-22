package properties


type Cluster_InstanceFleetConfig struct {
	
	
	
	
	
	Name interface{} `yaml:"Name,omitempty"`
	TargetOnDemandCapacity interface{} `yaml:"TargetOnDemandCapacity,omitempty"`
	TargetSpotCapacity interface{} `yaml:"TargetSpotCapacity,omitempty"`
	InstanceTypeConfigs interface{} `yaml:"InstanceTypeConfigs,omitempty"`
	LaunchSpecifications *Cluster_InstanceFleetProvisioningSpecifications `yaml:"LaunchSpecifications,omitempty"`
}

func (resource Cluster_InstanceFleetConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	return errs
}
