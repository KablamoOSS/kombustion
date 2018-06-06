package tasks

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// DeleteStack removes a cloudformation stack
func DeleteStack(stackName, profile, region string) {
	cf := GetCloudformationClient(profile, region)

	//See if the stack exists to begin with
	_, err := cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
	checkError(err)

	_, err = cf.DeleteStack(&cloudformation.DeleteStackInput{StackName: aws.String(stackName)})
	checkError(err)

	// status polling
	for {
		time.Sleep(2 * time.Second)
		status, _ := cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})

		if len(status.Stacks) > 0 {
			stackStatus := *status.Stacks[0].StackStatus
			fmt.Println(stackStatus)
			if stackStatus == cloudformation.StackStatusDeleteInProgress {
				continue
			}
		}
		break
	}

	// Make sure delete worked
	_, err = cf.DescribeStacks(&cloudformation.DescribeStacksInput{StackName: aws.String(stackName)})
	if err != nil {
		checkErrorDeletePoll(err)
	} else {
		fmt.Println("Delete successful")
		os.Exit(0)
	}
}
