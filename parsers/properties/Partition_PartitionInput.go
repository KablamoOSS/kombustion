package properties

	import "fmt"

type Partition_PartitionInput struct {
	
	
	
	Parameters interface{} `yaml:"Parameters,omitempty"`
	StorageDescriptor *Partition_StorageDescriptor `yaml:"StorageDescriptor,omitempty"`
	Values interface{} `yaml:"Values"`
}

func (resource Partition_PartitionInput) Validate() []error {
	errs := []error{}
	
	
	
	if resource.Values == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Values'"))
	}
	return errs
}
