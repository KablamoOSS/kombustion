package properties


type ReplicationGroup_NodeGroupConfiguration struct {
	
	
	
	
	PrimaryAvailabilityZone interface{} `yaml:"PrimaryAvailabilityZone,omitempty"`
	ReplicaCount interface{} `yaml:"ReplicaCount,omitempty"`
	Slots interface{} `yaml:"Slots,omitempty"`
	ReplicaAvailabilityZones interface{} `yaml:"ReplicaAvailabilityZones,omitempty"`
}

func (resource ReplicationGroup_NodeGroupConfiguration) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
