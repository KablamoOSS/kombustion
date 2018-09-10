package cloudformation

import (
	"github.com/aws/aws-sdk-go/aws"
	awsCF "github.com/aws/aws-sdk-go/service/cloudformation"
	"time"
)

type StackEventer interface {
	Open(string, string) string
	StackEvents(string) ([]*StackEvent, error)
}

type StackEvent struct {
	Time      time.Time
	Type      string
	LogicalID string
	Status    string
	Reason    string
}

type CloudFormationEventer struct {
	client *awsCF.CloudFormation
}

func NewEventer(client *awsCF.CloudFormation) *CloudFormationEventer {
	return &CloudFormationEventer{
		client: client,
	}
}

func (cfe *CloudFormationEventer) Open(profile, region string) string {
	acctID, cfClient := GetCloudformationClient(profile, region)
	cfe.client = cfClient
	return acctID
}

func (cfe *CloudFormationEventer) StackEvents(stack string) ([]*StackEvent, error) {
	events := make([]*StackEvent, 0)

	err := cfe.client.DescribeStackEventsPages(
		&awsCF.DescribeStackEventsInput{
			StackName: aws.String(stack),
		},
		func(page *awsCF.DescribeStackEventsOutput, lastPage bool) bool {
			for _, event := range page.StackEvents {

				stackEvent := &StackEvent{
					Time:      *event.Timestamp,
					Type:      *event.ResourceType,
					LogicalID: *event.LogicalResourceId,
					Status:    *event.ResourceStatus,
					Reason:    *event.ResourceStatusReason,
				}

				events = append(events, stackEvent)
			}

			return !lastPage
		},
	)

	if err != nil {
		return nil, err
	}

	return events, nil
}
