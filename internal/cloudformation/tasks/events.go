package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/ttacon/chalk"
)

var lineFormat = "%-19v | %-35v | %-40v | %-50v | %-30v "

// PrintStackEvents outputs the events of a stack
// TODO: add flags to allow printing all events, and default only to recent events
func PrintStackEvents(eventer cloudformation.StackEventer, stackName string) {
	events, err := eventer.StackEvents(stackName)
	checkError(err)

	printer.Step(fmt.Sprintf("Events for %s:", stackName))

	PrintStackEventHeader()

	for i, event := range events {
		var isLast bool
		if i+1 == len(events) {
			isLast = true
		}
		PrintStackEvent(event, isLast)
	}
}

// PrintStackEventHeader prints the header
func PrintStackEventHeader() {
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
func PrintStackEvent(event *cloudformation.StackEvent, isLast bool) {
	timeStamp := event.Time.Format("2006-01-02 15:04:05")
	resourceStatus := event.Status
	resourceType := event.Type
	logicalResourceID := event.LogicalID
	resourceStatusReason := event.Reason

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
