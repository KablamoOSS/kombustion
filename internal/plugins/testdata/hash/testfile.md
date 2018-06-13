# What is AWS CloudFormation?

AWS CloudFormation is a service that helps you model and set up your Amazon Web
Services resources so that you can spend less time managing those resources and
more time focusing on your applications that run in AWS. You create a template
that describes all the AWS resources that you want (like Amazon EC2 instances or
Amazon RDS DB instances), and AWS CloudFormation takes care of provisioning and
configuring those resources for you. You don't need to individually create and
configure AWS resources and figure out what's dependent on what; AWS
CloudFormation handles all of that. The following scenarios demonstrate how AWS
CloudFormation can help.

## Simplify Infrastructure Management

For a scalable web application that also includes a back-end database, you might
use an Auto Scaling group, an Elastic Load Balancing load balancer, and an
Amazon Relational Database Service database instance. Normally, you might use
each individual service to provision these resources. And after you create the
resources, you would have to configure them to work together. All these tasks
can add complexity and time before you even get your application up and running.

Instead, you can create or modify an existing AWS CloudFormation template. A
template describes all of your resources and their properties. When you use that
template to create an AWS CloudFormation stack, AWS CloudFormation provisions
the Auto Scaling group, load balancer, and database for you. After the stack has
been successfully created, your AWS resources are up and running. You can delete
the stack just as easily, which deletes all the resources in the stack. By using
AWS CloudFormation, you easily manage a collection of resources as a single
unit.
