# Quick start

Download the latest release from the [Installation](installation.md) page.

## Running the application

After you have downloaded your binary and ensured it exists in your PATH, ensure you can execute it:

```sh
kombustion -v
```

## Writing your template

Kombustion is fully backwards compatible with native CloudFormation templates. You can use any of your existing templates to get started, otherwise here's one to help you get started:

```example.yaml
AWSTemplateFormatVersion: 2010-09-09
Description: "A simple SNS topic"
Resources:
    MyTopic:
        Type: "AWS::SNS::Topic"
        Properties:
            TopicName: "MyExampleTopic"
```

## Ensuring your credentials are available

Kombustion will use the AWS CLI provided credentials location. If you already have the AWS CLI installed and have run an `aws configure`, you're good to go. If not, we can craft this file ourselves without the need to install the AWS CLI tooling. The default file location is ~/.aws/credentials and looks like the following:

```~/.aws/credentials
[default]
aws_access_key_id = AKIAJHNXGTMWJEXAMPLE
aws_secret_access_key = 5H34zbSHdwug3fQfylaq8mQ+3NaDJ4EXAMPLE
```

## Upserting your stack

The following will upsert the stack you have created when you specify the filename of the stack:

```sh
kombustion cf upsert example.yaml --stackName example-stack
```
