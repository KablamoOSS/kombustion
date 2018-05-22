# Generation

Kombustion creates the `/parsers` dynamically using a code generator. The generation application is a discrete binary that will use the [CloudFormation Specification](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-resource-specification.html) files to create the parsers that describe all CloudFormation resources, properties and outputs. With this, all CloudFormation-native resources are automatically supported by Kombustion.

## Updating and re-generating

If you would like to update to the latest version of the CloudFormation Specification and re-generate your parsers, you can execute the following and then rebuild after you have completed [initialization](initialization.md):

```sh
rm -rf parsers/ pluginParsers/ generate/source/
generate/generate
generate/generate pluginParsers
```
