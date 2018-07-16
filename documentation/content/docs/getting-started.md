+++
title = "Getting Started"
description = "How to start using Kombustion"
date = "2018-06-14T00:00:00+10:00"
weight = 20
draft = false
bref = "How to download and start using Kombustion"
toc = true
layout = "docs"
+++

## Install

Kombustion is built for Linux, FreeBSD and MacOS.

Get the latest release from the
[downloads page](/docs/downloads).

After downloading you will need to move the `kombustion` binary into your `$PATH`, and make it executable.

```bash
$ sudo chmod +x kombustion
$ sudo mv kombustion /usr/local/bin/kombustion
```

You can run `kombustion -v` to confirm you have the correct version installed.

## Create your project

In the root directory of your repository containing your CloudFormation templates run the following to create the `kombustion.yaml` file:

```bash
# Create a new kombustion.yaml file in the current directory
$ kombustion init
```

## Add a plugin

To use a plugin you first need to add it to your project

## Writing your template

Kombustion is fully backwards compatible with native CloudFormation templates. You can use any of your existing templates to get started, otherwise here's one to help you get started:

```yaml
# example.yaml
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

```
# ~/.aws/credentials
[default]
aws_access_key_id = AKIAXXXXXXXXXXXXXX
aws_secret_access_key = xxxxxxxxxxxxxxxxxxxxxxxxxx
```

## Upserting your stack

The following will upsert the stack you have created when you specify the filename of the stack:

```bash
kombustion upsert example.yaml
```
