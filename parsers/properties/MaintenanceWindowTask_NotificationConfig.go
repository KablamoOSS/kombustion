package properties

	import "fmt"

type MaintenanceWindowTask_NotificationConfig struct {
	
	
	
	NotificationArn interface{} `yaml:"NotificationArn"`
	NotificationType interface{} `yaml:"NotificationType,omitempty"`
	NotificationEvents interface{} `yaml:"NotificationEvents,omitempty"`
}

func (resource MaintenanceWindowTask_NotificationConfig) Validate() []error {
	errs := []error{}
	
	
	
	if resource.NotificationArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'NotificationArn'"))
	}
	return errs
}
