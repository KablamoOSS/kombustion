package tasks

var sampleKombYaml = `---
Name: Kombustion
Region: ""
Environments:
  ci:
    Parameters:
      BucketName: fooBucket
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
