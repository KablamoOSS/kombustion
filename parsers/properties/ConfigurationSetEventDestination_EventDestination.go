package properties

	import "fmt"

type ConfigurationSetEventDestination_EventDestination struct {
	
	
	
	
	
	Enabled interface{} `yaml:"Enabled,omitempty"`
	Name interface{} `yaml:"Name,omitempty"`
	MatchingEventTypes interface{} `yaml:"MatchingEventTypes"`
	KinesisFirehoseDestination *ConfigurationSetEventDestination_KinesisFirehoseDestination `yaml:"KinesisFirehoseDestination,omitempty"`
	CloudWatchDestination *ConfigurationSetEventDestination_CloudWatchDestination `yaml:"CloudWatchDestination,omitempty"`
}

func (resource ConfigurationSetEventDestination_EventDestination) Validate() []error {
	errs := []error{}
	
	
	
	
	
	if resource.MatchingEventTypes == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'MatchingEventTypes'"))
	}
	return errs
}
