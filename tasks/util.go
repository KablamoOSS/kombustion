package tasks

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/urfave/cli"
)

func checkError(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			log.Warn("No updates are to be performed.")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			log.Warn("The stack does not exist.")
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}

func checkErrorDeletePoll(err error) {
	if err != nil {
		if strings.Contains(err.Error(), "No updates are to be performed") {
			log.Warn("No updates are to be performed.")
			os.Exit(0)
		} else if strings.Contains(err.Error(), "Stack with id") && strings.Contains(err.Error(), "does not exist") {
			os.Exit(0)
		} else {
			log.Fatal(err)
		}
	}
}

func getParamMap(c *cli.Context) map[string]string {
	paramMap := make(map[string]string)
	params := c.StringSlice("param")
	for _, param := range params {
		parts := strings.Split(param, "=")
		if len(parts) > 1 {
			paramMap[parts[0]] = strings.Join(parts[1:], "=")
		}
	}
	return paramMap
}

func printStackEvents(cf *awsCF.CloudFormation, stackName string) {
	status, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
	checkError(err)

	fmt.Println()
	fmt.Printf(" %-19v | %-22v | %-30v | %v | %v | \n", "Time", "Status", "Type", "LogicalID", "Status Reason")
	for _, event := range status.StackEvents {
		if event.Timestamp != nil {
			fmt.Printf(" %-19v |", event.Timestamp.Format("2006-01-2 15:04:05"))
		}
		if event.ResourceStatus != nil {
			fmt.Printf(" %-22v |", *event.ResourceStatus)
		}
		if event.ResourceType != nil {
			fmt.Printf(" %-30v |", *event.ResourceType)
		}
		if event.LogicalResourceId != nil {
			fmt.Printf(" %v |", *event.LogicalResourceId)
		}
		if event.ResourceStatusReason != nil {
			fmt.Printf(" %v |", *event.ResourceStatusReason)
		}
		fmt.Println()
	}
}
