package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/ttacon/chalk"
	"github.com/aws/aws-sdk-go/aws"
)

var lineFormat  = "%-19v | %-35v | %-40v | %-50v | %-30v "

// PrintStackEvents outputs the events of a stack
// TODO: add flags to allow printing all events, and default only to recent events
func PrintStackEvents(cf *awsCF.CloudFormation, stackName string) {
	status, err := cf.DescribeStackEvents(&awsCF.DescribeStackEventsInput{StackName: aws.String(stackName)})
	checkError(err)

	printer.Step(fmt.Sprintf("Events for %s:", stackName))

	PrintStackEventHeader()

	for i, event := range status.StackEvents {
		var isLast bool
		if i + 1 == len(status.StackEvents) {
			isLast = true
		}
		PrintStackEvent(event, isLast)
	}
}

// PrintStackEventHeader prints the header
func PrintStackEventHeader(){
	printer.SubStep(
		fmt.Sprintf(
			chalk.Bold.TextStyle(lineFormat),
			"Time",
			"Status",
			"Type",
			"LogicalID",
			"Status Reason",
		),
		1,
		false,
		true,
	)
}

// PrintStackEvent prints a single event
func PrintStackEvent(event *awsCF.StackEvent, isLast bool) {
	var timeStamp,
	resourceStatus,
	resourceType,
	logicalResourceID,
	resourceStatusReason string
	if event.Timestamp != nil {
		timeStamp = event.Timestamp.Format("2006-01-02 15:04:05")
	}
	if event.ResourceStatus != nil {
		resourceStatus= *event.ResourceStatus
	}
	if event.ResourceType != nil {
		resourceType	= *event.ResourceType
	}
	if event.LogicalResourceId != nil {
		logicalResourceID	= *event.LogicalResourceId
	}
	if event.ResourceStatusReason != nil {
		resourceStatusReason	= *event.ResourceStatusReason
	}

	printer.SubStep(
		fmt.Sprintf(
			lineFormat,
			timeStamp,
			resourceStatus,
			resourceType,
			logicalResourceID,
			resourceStatusReason,
		),
		1,
		isLast,
		true,
	)
}
