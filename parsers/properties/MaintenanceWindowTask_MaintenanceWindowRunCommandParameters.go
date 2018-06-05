package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type MaintenanceWindowTask_MaintenanceWindowRunCommandParameters struct {
	Comment            interface{}                               `yaml:"Comment,omitempty"`
	DocumentHash       interface{}                               `yaml:"DocumentHash,omitempty"`
	DocumentHashType   interface{}                               `yaml:"DocumentHashType,omitempty"`
	OutputS3BucketName interface{}                               `yaml:"OutputS3BucketName,omitempty"`
	OutputS3KeyPrefix  interface{}                               `yaml:"OutputS3KeyPrefix,omitempty"`
	Parameters         interface{}                               `yaml:"Parameters,omitempty"`
	ServiceRoleArn     interface{}                               `yaml:"ServiceRoleArn,omitempty"`
	TimeoutSeconds     interface{}                               `yaml:"TimeoutSeconds,omitempty"`
	NotificationConfig *MaintenanceWindowTask_NotificationConfig `yaml:"NotificationConfig,omitempty"`
}

func (resource MaintenanceWindowTask_MaintenanceWindowRunCommandParameters) Validate() []error {
	errs := []error{}

	return errs
}
