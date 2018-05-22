package properties


type Cluster_JobFlowInstancesConfig struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Ec2KeyName interface{} `yaml:"Ec2KeyName,omitempty"`
	Ec2SubnetId interface{} `yaml:"Ec2SubnetId,omitempty"`
	EmrManagedMasterSecurityGroup interface{} `yaml:"EmrManagedMasterSecurityGroup,omitempty"`
	EmrManagedSlaveSecurityGroup interface{} `yaml:"EmrManagedSlaveSecurityGroup,omitempty"`
	HadoopVersion interface{} `yaml:"HadoopVersion,omitempty"`
	ServiceAccessSecurityGroup interface{} `yaml:"ServiceAccessSecurityGroup,omitempty"`
	TerminationProtected interface{} `yaml:"TerminationProtected,omitempty"`
	Placement *Cluster_PlacementType `yaml:"Placement,omitempty"`
	AdditionalSlaveSecurityGroups interface{} `yaml:"AdditionalSlaveSecurityGroups,omitempty"`
	AdditionalMasterSecurityGroups interface{} `yaml:"AdditionalMasterSecurityGroups,omitempty"`
	CoreInstanceGroup *Cluster_InstanceGroupConfig `yaml:"CoreInstanceGroup,omitempty"`
	MasterInstanceGroup *Cluster_InstanceGroupConfig `yaml:"MasterInstanceGroup,omitempty"`
	CoreInstanceFleet *Cluster_InstanceFleetConfig `yaml:"CoreInstanceFleet,omitempty"`
	MasterInstanceFleet *Cluster_InstanceFleetConfig `yaml:"MasterInstanceFleet,omitempty"`
}

func (resource Cluster_JobFlowInstancesConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	return errs
}
