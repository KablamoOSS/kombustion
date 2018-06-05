package tasks

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

var Delete_Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "region, r",
		Usage: "region to delete from",
		Value: "ap-southeast-2",
	},
}

func Delete(c *cli.Context) {
	deleteStack(c, getCF(c.GlobalString("profile"), c.String("region")))
}

func deleteStack(c *cli.Context, cf *cloudformation.CloudFormation) {
	stackName := c.Args().Get(0)

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
