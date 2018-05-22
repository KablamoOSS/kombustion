package properties

	import "fmt"

type RecordSetGroup_RecordSet struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Comment interface{} `yaml:"Comment,omitempty"`
	Failover interface{} `yaml:"Failover,omitempty"`
	HealthCheckId interface{} `yaml:"HealthCheckId,omitempty"`
	HostedZoneId interface{} `yaml:"HostedZoneId,omitempty"`
	HostedZoneName interface{} `yaml:"HostedZoneName,omitempty"`
	Name interface{} `yaml:"Name"`
	Region interface{} `yaml:"Region,omitempty"`
	SetIdentifier interface{} `yaml:"SetIdentifier,omitempty"`
	TTL interface{} `yaml:"TTL,omitempty"`
	Type interface{} `yaml:"Type"`
	Weight interface{} `yaml:"Weight,omitempty"`
	ResourceRecords interface{} `yaml:"ResourceRecords,omitempty"`
	GeoLocation *RecordSetGroup_GeoLocation `yaml:"GeoLocation,omitempty"`
	AliasTarget *RecordSetGroup_AliasTarget `yaml:"AliasTarget,omitempty"`
}

func (resource RecordSetGroup_RecordSet) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
