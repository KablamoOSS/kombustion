# Environmental variables

In many cases, you may want to use the same templates through all your
deployment environments to keep consistency during your build process. However,
some values may change between environments, such as the endpoints your
applications use or parameters used for testing.

Kombustion facilitates this by providing the ability to template values and
replace them based on the environment you specify. Below is an example of a
template that has a replaceable variable:

```example.yaml
AWSTemplateFormatVersion: 2010-09-09
Description: "A simple SNS topic"
Resources:
    MyTopic:
        Type: "AWS::SNS::Topic"
        Properties:
            TopicName: "My{{topictype}}Topic"
```

In the `examples/environment.yaml` file, we define what value this is for
specific environments and also provide a default value for the key if no other
match is found, under the `_default` section:

```examples/environment.yaml
_default:
    topictype: Yellow
dev:
    topictype: Green
test:
    topictype: Blue
preprod:
    topictype: Red
prod:
    topictype: Red
```

The resulting template when the environment is `test`, will be:

```compiled/example.yaml
AWSTemplateFormatVersion: 2010-09-09
Description: "A simple SNS topic"
Resources:
    MyTopic:
        Type: "AWS::SNS::Topic"
        Properties:
            TopicName: "MyBlueTopic"
```
