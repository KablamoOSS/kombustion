package properties


type MaintenanceWindowTask_TaskInvocationParameters struct {
	
	
	
	
	MaintenanceWindowStepFunctionsParameters *MaintenanceWindowTask_MaintenanceWindowStepFunctionsParameters `yaml:"MaintenanceWindowStepFunctionsParameters,omitempty"`
	MaintenanceWindowRunCommandParameters *MaintenanceWindowTask_MaintenanceWindowRunCommandParameters `yaml:"MaintenanceWindowRunCommandParameters,omitempty"`
	MaintenanceWindowLambdaParameters *MaintenanceWindowTask_MaintenanceWindowLambdaParameters `yaml:"MaintenanceWindowLambdaParameters,omitempty"`
	MaintenanceWindowAutomationParameters *MaintenanceWindowTask_MaintenanceWindowAutomationParameters `yaml:"MaintenanceWindowAutomationParameters,omitempty"`
}

func (resource MaintenanceWindowTask_TaskInvocationParameters) Validate() []error {
	errs := []error{}
	
	
	
	
	return errs
}
