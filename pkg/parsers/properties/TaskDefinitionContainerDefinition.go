package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionContainerDefinition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
type TaskDefinitionContainerDefinition struct {
	Cpu                    interface{}                     `yaml:"Cpu,omitempty"`
	DisableNetworking      interface{}                     `yaml:"DisableNetworking,omitempty"`
	Essential              interface{}                     `yaml:"Essential,omitempty"`
	Hostname               interface{}                     `yaml:"Hostname,omitempty"`
	Image                  interface{}                     `yaml:"Image,omitempty"`
	Memory                 interface{}                     `yaml:"Memory,omitempty"`
	MemoryReservation      interface{}                     `yaml:"MemoryReservation,omitempty"`
	Name                   interface{}                     `yaml:"Name,omitempty"`
	Privileged             interface{}                     `yaml:"Privileged,omitempty"`
	ReadonlyRootFilesystem interface{}                     `yaml:"ReadonlyRootFilesystem,omitempty"`
	User                   interface{}                     `yaml:"User,omitempty"`
	WorkingDirectory       interface{}                     `yaml:"WorkingDirectory,omitempty"`
	DockerLabels           interface{}                     `yaml:"DockerLabels,omitempty"`
	LogConfiguration       *TaskDefinitionLogConfiguration `yaml:"LogConfiguration,omitempty"`
	Command                interface{}                     `yaml:"Command,omitempty"`
	DnsServers             interface{}                     `yaml:"DnsServers,omitempty"`
	DockerSecurityOptions  interface{}                     `yaml:"DockerSecurityOptions,omitempty"`
	EntryPoint             interface{}                     `yaml:"EntryPoint,omitempty"`
	Environment            interface{}                     `yaml:"Environment,omitempty"`
	ExtraHosts             interface{}                     `yaml:"ExtraHosts,omitempty"`
	Links                  interface{}                     `yaml:"Links,omitempty"`
	Ulimits                interface{}                     `yaml:"Ulimits,omitempty"`
	DnsSearchDomains       interface{}                     `yaml:"DnsSearchDomains,omitempty"`
	MountPoints            interface{}                     `yaml:"MountPoints,omitempty"`
	PortMappings           interface{}                     `yaml:"PortMappings,omitempty"`
	VolumesFrom            interface{}                     `yaml:"VolumesFrom,omitempty"`
	LinuxParameters        *TaskDefinitionLinuxParameters  `yaml:"LinuxParameters,omitempty"`
}

func (resource TaskDefinitionContainerDefinition) Validate() []error {
	errs := []error{}

	return errs
}
