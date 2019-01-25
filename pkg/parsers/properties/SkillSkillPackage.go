package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SkillSkillPackage Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ask-skill-skillpackage.html
type SkillSkillPackage struct {
	S3Bucket        interface{} `yaml:"S3Bucket"`
	S3BucketRole    interface{} `yaml:"S3BucketRole,omitempty"`
	S3Key           interface{} `yaml:"S3Key"`
	S3ObjectVersion interface{} `yaml:"S3ObjectVersion,omitempty"`
	Overrides       interface{} `yaml:"Overrides,omitempty"`
}

// SkillSkillPackage validation
func (resource SkillSkillPackage) Validate() []error {
	errors := []error{}

	return errors
}
