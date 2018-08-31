package tasks

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/KablamoOSS/kombustion/internal/coretest"
)

var sampleKombYaml = `---
Name: Kombustion
Region: ""
Environments:
  development:
    Parameters:
      Environment: development
GenerateDefaultOutputs: false
Tags: {}
`

var sampleKombLock = `plugins: {}`

var sampleYaml = `AWSTemplateFormatVersion: 2010-09-09
Description: S3 test bucket
Parameters:
  BucketName: 
    Type: "String"
    Default: "testBucket"
    Description: "S3 bucket name"

Mappings: {}
Resources:
  testBucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Ref: BucketName
      AccessControl: PublicRead
      Tags:
      - Key: Name
        Value: 123
`

// TODO: This is super fragile, but works as a quick and dirty first pass to
// some testing over the generate task. We should really decompose this down to
// specific properties we want to test for.
var expectedOutput = `AWSTemplateFormatVersion: "2010-09-09"
Description: S3 test bucket
Parameters:
  BucketName:
    Default: testBucket
    Description: S3 bucket name
    Type: String
Resources:
  testBucket:
    Type: AWS::S3::Bucket
    Properties:
      AccessControl: PublicRead
      BucketName: BucketName
      Tags:
      - Key: Name
        Value: 123
`

func TestSimpleGenerate(t *testing.T) {
	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombYaml), "kombustion.yaml")
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")
	objectStore.Put([]byte(sampleYaml), "test.yaml")

	assert.NotPanics(
		t,
		func() {
			generate(
				objectStore,
				"test.yaml",
				map[string]string{},
				"",
				"",
				"compiled",
				true,
				"ci",
				false,
			)
		},
	)

	output, _ := objectStore.Get("compiled", "test.yaml")
	assert.Equal(t, expectedOutput, string(output))
}
