## Kombustion

[![Build Status](https://travis-ci.org/KablamoOSS/kombustion.svg?branch=master)](https://travis-ci.org/KablamoOSS/kombustion)
[![](https://tokei.rs/b1/github/KablamoOSS/kombustion)](https://github.com/KablamoOSS/kombustion)
[![Go Report Card](https://goreportcard.com/badge/github.com/KablamoOSS/kombustion)](https://goreportcard.com/report/github.com/KablamoOSS/kombustion)
[![Coverage Status](https://coveralls.io/repos/github/KablamoOSS/kombustion/badge.svg?branch=master)](https://coveralls.io/github/KablamoOSS/kombustion?branch=master)

> Extend CloudFormation with plugins

Kombustion uses plugins to preprocess and extend your CloudFormation templates.

In addition to generating templates, Kombustion can also create, update and
delete your CloudFormation stacks.

Kombustion has automatic support for new CloudFormation types as they are
released.

See the [Quick start](https://kombustion.io/docs/getting-started/) for more details.

## Getting Started

Kombustion is built for Linux, FreeBSD, MacOS and Windows.

Get the latest release from the
[release page](https://github.com/KablamoOSS/kombustion/releases).

After downloading for MacOS or Linux, you will need to move the `kombustion`
binary into your `$PATH`, and make it executable.

```bash
sudo chmod +x kombustion
sudo cp kombustion /usr/local/bin/kombustion
```

## Usage

Initialise a `kombustion.yaml` file with the following:

```sh
$ kombustion init
```

This is an example of `kombustion.yaml`, it should be committed to version control.

```yaml
# Name of this project. This is used with `--environment` to create a stack name, which can
# be overridden with `--stack-name`
Name: KombustionExample
# Region is the default region stacks will be deployed into. Can be overridden with `--region us-east-2`
Region: us-east-1
# Plugins can be installed by running `kombustion add github.com/KablamoOSS/kombustion-plugin-serverless`
Plugins:
  github.com/KablamoOSS/kombustion-plugin-serverless@0.1.0:
    Name: github.com/KablamoOSS/kombustion-plugin-serverless
    Version: 0.1.0
Environments:
  Development:
    # Optionally whitelist the accounts for the environment `development`, this will
    # ensure CloudFormation actions are only performed in this account
    AccountIDs:
    - "1234567890"
    # Parameters are added to the CloudFormation Stack during upsert
    Parameters:
      Environment: development
# Kombustion can generate default outputs for your resources to use as references in other
# stacks.
GenerateDefaultOutputs: false
```

### CloudFormation Stack Management

Upsert a CloudFormation template:

```sh
$ kombustion upsert examples/stacks/test.yaml --stackName test-stack
```

Delete a CloudFormation stack:

```sh
$ kombustion delete examples/stacks/test.yaml
```

Print all the events for a stack:

```sh
$ kombustion events examples/stacks/test.yaml
```

#### StackName

You don't need to specify `--stack-name`, instead when you pass an environment `--environment` it
gets merged in with the project name from `kombustion.yaml` and the filename as `{ProjectName}-{FileName}-{Environment}`.

This applies to `upsert`,`delete`, and `events`.

### CloudFormation Stacks

A stack template is written in the same way as standard CloudFormation.
Kombustion allows plugins to extend the syntax, but the end result is always
standard CloudFormation.

The following example shows how a small definition for a
[bastion host](https://en.wikipedia.org/wiki/Bastion_host), can be processed
into a bigger template. This lets your plugin maintain safe, sane defaults, and
ensure you don't miss any required fields.

```yaml
# In this example we're going to create a bastion host.
# This is a small EC2 instance, configured with a public IP
# and a security group to allow us to SSH into our AWS cloud.
AWSTemplateFormatVersion: 2010-09-09
Description: Example EC2 Instance
Parameters: {}
Mappings: {}
Resources:
  BastionHost:
    Type: Kombustion::Examples::BastionHost
    Properties:
      # In this example, this key would have been uploaded to AWS
      KeyName: my-ssh-key
      Size: t2.micro
      # Using a filter, find the most recent AMI of Amazon Linux 2
      AmiFilter:
        VirtualizationType: "hvm"
        Name: "amzn2-ami-*",
        RootDeviceType: "ebs"
        owners: ["amazon"],
        Latest: true
```

The Plugin `Kombustion::Examples::BastionHost` is used to generate the following
template. It uses the AmiFilter to find the correct AMI, and creates two
parameters for the `KeyName` and `SSHLocation`. The latter being the IP address
allowed through the security group.

```yaml
AWSTemplateFormatVersion: 2010-09-09
Description: Example EC2 Instance
Parameters:
  KombustionExampleBastionHostKeyName:
    Description: Name of an existing EC2 KeyPair to enable SSH access to the instances
    Type: 'AWS::EC2::KeyPair::KeyName'
    Default: 'my-ssh-key'
    ConstraintDescription: must be the name of an existing EC2 KeyPair.
  KombustionExampleBastionHostSSHLocation:
    Description: The IP address range that can be used to SSH to the EC2 instances
    Type: String
    MinLength: '9'
    MaxLength: '18'
    Default: 0.0.0.0/0
    AllowedPattern: '(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})/(\d{1,2})'
    ConstraintDescription: must be a valid IP CIDR range of the form x.x.x.x/x.
Mappings: {}
Resources:
Resources:
  EC2Instance:
    Type: 'AWS::EC2::Instance'
    Properties:
      InstanceType: !Ref InstanceType
      SecurityGroups:
        - !Ref InstanceSecurityGroup
      KeyName: !Ref KombustionExampleBastionHostKeyName
      ImageId: 'ami-c267b0a0'
  InstanceSecurityGroup:
    Type: 'AWS::EC2::SecurityGroup'
    Properties:
      GroupDescription: Enable SSH access
      SecurityGroupIngress:
        - IpProtocol: tcp
          FromPort: '22'
          ToPort: '22'
          CidrIp: !Ref KombustionExampleBastionHostSSHLocation
  IPAddress:
    Type: 'AWS::EC2::EIP'
  IPAssoc:
    Type: 'AWS::EC2::EIPAssociation'
    Properties:
      InstanceId: !Ref EC2Instance
      EIP: !Ref IPAddress
```

Check out the
[examples](https://github.com/KablamoOSS/Kombustion/tree/master/examples/)
directory for example stacks.


#### Credentials

Kombustion uses the same method as the `aws` cli to get
[credential information](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).
You can either use the standard environment variables `AWS_ACCESS_KEY_ID`,
`AWS_SECRET_ACCESS_KEY`, and `AWS_SESSION_TOKEN`.

Or use a profile you have configured, for example:

```sh
$ kombustion --profile myAwsProfile upsert examples/stacks/test.yaml --stackName test-stack
```

### Plugins

> Kombustion plugins are not yet supported on Windows, due to
> [this issue](https://github.com/golang/go/issues/19282). Please use Docker or
> WSL in the meantime.

Install a plugin:

```sh
$ kombustion add github.com/Example/ExamplePlugin
```

## Contributing

Please read
[CONTRIBUTING.md](https://github.com/KablamoOSS/kombustion/blob/master/CONTRIBUTING.md)
for details on our code of conduct, and the process for submitting pull requests
to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available,
see the
[tags on this repository](https://github.com/KablamoOSS/kombustion/tags).

## Maintainers

Kombustion is primarily maintained by the [Kablamo](https://www.kablamo.com.au/)
team. Pull requests are welcome.

## Acknowledgements

The Kombustion logo is based on an original design by Renee French.

## License

This project is licensed under the
[MIT License](https://github.com/KablamoOSS/kombustion/blob/master/LICENSE).

---

Made with :heart: in Australia.
