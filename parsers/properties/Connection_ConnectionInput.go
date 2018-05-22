package properties

	import "fmt"

type Connection_ConnectionInput struct {
	
	
	
	
	
	
	ConnectionProperties interface{} `yaml:"ConnectionProperties"`
	ConnectionType interface{} `yaml:"ConnectionType"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	PhysicalConnectionRequirements *Connection_PhysicalConnectionRequirements `yaml:"PhysicalConnectionRequirements,omitempty"`
	MatchCriteria interface{} `yaml:"MatchCriteria,omitempty"`
}

func (resource Connection_ConnectionInput) Validate() []error {
	errs := []error{}
	
	
	
	
	
	
	if resource.ConnectionProperties == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ConnectionProperties'"))
	}
	if resource.ConnectionType == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'ConnectionType'"))
	}
	return errs
}
