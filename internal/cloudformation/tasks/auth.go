package tasks

import (
	"fmt"
	"os"
	"strings"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// GetCloudformationClient from the standard credential chain
func GetCloudformationClient(profile string, region string) *cloudformation.CloudFormation {
	var awsConfig *aws.Config

	if region == "" {
		printer.Fatal(
			fmt.Errorf("No region has been set"),
			fmt.Sprintf(
				"Pass a region in via the cli with `--region us-east-1` or set a default `region` in kombustion.yaml",
			),
			"http://kombustion.io/manifest#region",
		)
	}

	// If we have a custom env var for assuming a role, lets assume
	if len(os.Getenv("ASSUMED_ROLE")) > 0 {
		// TODO: Check for all env params, and err if not all exist
		creds := getAWSCredentials()
		awsConfig = &aws.Config{Credentials: creds, Region: aws.String(region)}
	} else {
		awsConfig = &aws.Config{Region: aws.String(region)}
	}

	awsSession := getSession(profile)
	cf := cloudformation.New(awsSession, awsConfig)
	return cf
}

func getAWSCredentials() *credentials.Credentials {
	assumedRole := os.Getenv("ASSUMED_ROLE")
	mfaSerial := os.Getenv("MFA_SERIAL")
	awsMfaToken := os.Getenv("TOKEN")
	sess := must(session.NewSession())

	if len(mfaSerial) > 0 {
		return stscreds.NewCredentials(sess, assumedRole, func(p *stscreds.AssumeRoleProvider) {
			p.SerialNumber = aws.String(mfaSerial)
			p.TokenCode = aws.String(awsMfaToken)
			p.Duration = 3600
		})
	}

	return stscreds.NewCredentials(sess, assumedRole)
}

// Get a session, or fatal error out explaining why we didn't get one
func getSession(profile string) *session.Session {
	var options session.Options

	options = session.Options{
		// We pass a custom token provider here
		// to ensure we can stop the printer while we wait for
		// the mfa token
		AssumeRoleTokenProvider: mfaTokenProvider,

		Config: aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}

	if len(profile) > 0 {
		options.Profile = profile
		options.SharedConfigState = session.SharedConfigEnable
	}
	awsSession := must(session.NewSessionWithOptions(options))

	return awsSession
}

// Custom token provider to ensure we can stop the printer while we wait for
// the mfa token
func mfaTokenProvider() (string, error) {
	var v string
	printer.Stop()
	fmt.Printf("Enter MFA code: ")
	_, err := fmt.Scanln(&v)
	return v, err
}

// Ensure a session is returned, else fatal with an error explaning why no session was found
func must(sess *session.Session, err error) *session.Session {

	// if err != nil {
	// 	if awsErr, ok := err.(awserr.Error); ok {
	// 		// Get error details
	// 		log.Println("Error:", awsErr.Code(), awsErr.Message())

	// 		// Prints out full error message, including original error if there was one.
	// 		log.Println("Error:", awsErr.Error())

	// 		// Get original error
	// 		if origErr := awsErr.OrigErr(); origErr != nil {
	// 			// operate on original error.
	// 		}
	// 	} else {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	if err != nil {
		if strings.Contains(err.Error(), "NoCredentialProviders") {
			printer.Fatal(
				err,
				"You need to provide access credentials to your AWS account.",
				"",
			)
		}

		printer.Fatal(
			err,
			"",
			"",
		)
	}
	return sess
}
