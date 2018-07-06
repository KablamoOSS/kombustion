package tasks

import (
	"fmt"
	"os"
	"time"

	printer "github.com/KablamoOSS/go-cli-printer"

	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
)

// UpsertStack -
func UpsertStack(
	templateBody []byte,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
) {

	var err error
	var action string

	// use template from file
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		action = "Updating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	} else {
		action = "Creating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateBody: aws.String(string(templateBody)),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	}
	checkError(err)

	processUpsert(stackName, action, cf)
}

// UpsertStackViaS3 -
func UpsertStackViaS3(
	templateURL string,
	parameters []*awsCF.Parameter,
	capabilities []*string,
	stackName string,
	cf *awsCF.CloudFormation,
) {

	var err error
	var action string

	// use cf template url
	_, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
	if err == nil { //update
		action = "Updating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))

		_, err = cf.UpdateStack(&awsCF.UpdateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	} else { //create
		action = "Creating"
		printer.Step(fmt.Sprintf("%s Stack %s:", action, stackName))
		_, err = cf.CreateStack(&awsCF.CreateStackInput{
			StackName:    aws.String(stackName),
			TemplateURL:  aws.String(templateURL),
			Parameters:   parameters,
			Capabilities: capabilities,
		})
	}
	checkError(err)

	processUpsert(stackName, action, cf)

}

func processUpsert(stackName, action string, cf *awsCF.CloudFormation) {
	var err error
	var status *awsCF.DescribeStacksOutput

	PrintStackEventHeader()
	for {
		printer.Progress(action)
		time.Sleep(2 * time.Second)
		status, err = cf.DescribeStacks(&awsCF.DescribeStacksInput{StackName: aws.String(stackName)})
		checkError(err)

		events, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
		checkError(err)

		if len(status.Stacks) > 0 {

			stack := status.Stacks[0]
			stackStatus := *stack.StackStatus

			if len(events.StackEvents) > 0 {
				PrintStackEvent(events.StackEvents[0], false)
			}

			if stackStatus != awsCF.StackStatusCreateInProgress &&
				stackStatus != awsCF.StackStatusUpdateInProgress &&
				stackStatus != awsCF.StackStatusUpdateCompleteCleanupInProgress {
				if stackStatus == awsCF.StackStatusCreateComplete ||
					stackStatus == awsCF.StackStatusUpdateComplete {
					printer.SubStep(
						fmt.Sprintf("Success %s Stack %s", action, stackName),
						1,
						true,
						true,
					)
					os.Exit(0)
				} else {
					printer.SubStep(
						fmt.Sprintf("Fail %s Stack %s", action, stackName),
						1,
						true,
						true,
					)

					printer.Error(fmt.Errorf("Upsert Failed"), "", "")
					time.Sleep(2 * time.Second)
					os.Exit(1)
				}
			}
		}
	}
}
