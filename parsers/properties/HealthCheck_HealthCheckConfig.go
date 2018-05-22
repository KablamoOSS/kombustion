package properties

	import "fmt"

type HealthCheck_HealthCheckConfig struct {
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	EnableSNI interface{} `yaml:"EnableSNI,omitempty"`
	FailureThreshold interface{} `yaml:"FailureThreshold,omitempty"`
	FullyQualifiedDomainName interface{} `yaml:"FullyQualifiedDomainName,omitempty"`
	HealthThreshold interface{} `yaml:"HealthThreshold,omitempty"`
	IPAddress interface{} `yaml:"IPAddress,omitempty"`
	InsufficientDataHealthStatus interface{} `yaml:"InsufficientDataHealthStatus,omitempty"`
	Inverted interface{} `yaml:"Inverted,omitempty"`
	MeasureLatency interface{} `yaml:"MeasureLatency,omitempty"`
	Port interface{} `yaml:"Port,omitempty"`
	RequestInterval interface{} `yaml:"RequestInterval,omitempty"`
	ResourcePath interface{} `yaml:"ResourcePath,omitempty"`
	SearchString interface{} `yaml:"SearchString,omitempty"`
	Type interface{} `yaml:"Type"`
	ChildHealthChecks interface{} `yaml:"ChildHealthChecks,omitempty"`
	Regions interface{} `yaml:"Regions,omitempty"`
	AlarmIdentifier *HealthCheck_AlarmIdentifier `yaml:"AlarmIdentifier,omitempty"`
}

func (resource HealthCheck_HealthCheckConfig) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	if resource.Type == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Type'"))
	}
	return errs
}
