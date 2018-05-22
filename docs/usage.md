# Usage

## Basic Usage

Generate a CloudFormation template (from `./configs/test.yaml`):

```sh
kombustion cf generate configs/test.yaml && cat compiled/test.yaml
```

Upsert a CloudFormation template:

```sh
kombustion cf upsert configs/test.yaml --stackName test-stack
```

Delete a CloudFormation stack:

```sh
kombustion cf delete configs/test.yaml
```

Print all the events for a stack:

```sh
kombustion cf events configs/test.yaml
```

## Plugin management

!> Kombustion plugins are not yet supported on Windows. Please use Docker or WSL in the meantime.

Install a plugin:

```sh
kombustion cf plugins get mypluginname
```

List all installed plugins:

```sh
kombustion cf plugins list
```

Delete an installed plugin:

```sh
kombustion cf plugins delete mypluginname
```

## IAM arguments

Using Roles and MFA:

```sh
  TOKEN=000000 \
  MFA_SERIAL=arn:aws:iam::123456789012:mfa/stackCreator \
  ASSUMED_ROLE=arn:aws:iam::123456789012:role/god \
  kombustion cf upsert configs/test.yaml --stackName test-stack
```

Using a profile on in your `~/.aws/credentials`:

```sh
  kombustion cf --profile=MyProfile upsert configs/test.yaml
```
