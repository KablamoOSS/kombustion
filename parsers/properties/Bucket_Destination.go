package properties

	import "fmt"

type Bucket_Destination struct {
	
	
	
	
	BucketAccountId interface{} `yaml:"BucketAccountId,omitempty"`
	BucketArn interface{} `yaml:"BucketArn"`
	Format interface{} `yaml:"Format"`
	Prefix interface{} `yaml:"Prefix,omitempty"`
}

func (resource Bucket_Destination) Validate() []error {
	errs := []error{}
	
	
	
	
	if resource.BucketArn == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'BucketArn'"))
	}
	if resource.Format == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Format'"))
	}
	return errs
}
