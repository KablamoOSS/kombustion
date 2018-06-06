package tasks

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// GetCloudformationClient from the standard credential chain
func GetCloudformationClient(profile string, region string) *cloudformation.CloudFormation {
	var awsConfig *aws.Config

	// If we have a custom env var for assuming a role, lets assume
	if len(os.Getenv("ASSUMED_ROLE")) > 0 {
		// TODO: Check for all env params, and err if not all exist
		creds := getAWSCredentials()
		awsConfig = &aws.Config{Credentials: creds, Region: aws.String(region)}
	} else {
		awsConfig = &aws.Config{Region: aws.String(region)}
	}

	awsSession := session.Must(getSession(profile))
	cf := cloudformation.New(awsSession, awsConfig)
	return cf
}

func getAWSCredentials() *credentials.Credentials {
	assumedRole := os.Getenv("ASSUMED_ROLE")
	mfaSerial := os.Getenv("MFA_SERIAL")
	awsMfaToken := os.Getenv("TOKEN")
	sess := session.Must(session.NewSession())

	if len(mfaSerial) > 0 {
		return stscreds.NewCredentials(sess, assumedRole, func(p *stscreds.AssumeRoleProvider) {
			p.SerialNumber = aws.String(mfaSerial)
			p.TokenCode = aws.String(awsMfaToken)
			p.Duration = 3600
		})
	}

	return stscreds.NewCredentials(sess, assumedRole)
}

func getSession(profile string) (*session.Session, error) {
	var options session.Options

	options = session.Options{
		AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	}

	if len(profile) > 0 {
		options.Profile = profile
		options.SharedConfigState = session.SharedConfigEnable
	}
	awsSession, err := session.NewSessionWithOptions(options)

	return awsSession, err
}
