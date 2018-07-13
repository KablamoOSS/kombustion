package manifest

import (
	"fmt"
	"strings"

	printer "github.com/KablamoOSS/go-cli-printer"
	"gopkg.in/AlecAivazis/survey.v1"
)

// InitaliseNewManifest creates a new manifest with a survey
func InitaliseNewManifest() error {
	// TODO: Check if there is a manifest file and exit

	// Load the manifest file from this directory
	manifestExists := CheckManifestExists()
	if manifestExists {
		printer.Fatal(
			fmt.Errorf("Sorry we can't create a new kombustion.yaml, one already exists."),
			"If you want to re-initialise your kombustion.yaml file, first remove it.",
			"https://www.kombustion.io/api/manifest/",
		)
	}

	// Survey the user for required info
	name, environments, err := surveyForInitialManifest()
	if err != nil {
		return err
	}

	manifest := Manifest{
		Name:         name,
		Environments: environments,
	}

	err = WriteManifestToDisk(&manifest)
	if err != nil {
		return err
	}
	return nil
}

// surveyForInitialManifest - Prompt the user to fill out the required fields,
// but check if we have a flag for them
func surveyForInitialManifest() (
	name string,
	environments map[string]Environment,
	err error,
) {

	// name
	surveyName, err := surveyForName()
	if err != nil {
		return name, environments, err
	}
	name = surveyName

	//environments
	surveyEnvironments, err := surveyForEnvironments()
	if err != nil {
		return name, environments, err
	}
	environments = surveyEnvironments

	return name, environments, nil
}

// surveyForName - Prompt for the name of the project
func surveyForName() (string, error) {
	// the questions to ask
	var surveyQuestions = []*survey.Question{
		{
			Name:     "Name",
			Prompt:   &survey.Input{Message: "What is the name of this project?"},
			Validate: survey.Required,
		},
	}

	// the answers will be written to this struct
	surveyAnswers := struct {
		Name string // survey will match the question and field names
	}{}

	// perform the questions
	err := survey.Ask(surveyQuestions, &surveyAnswers)
	if err != nil {
		return "", err
	}

	// TODO: Add a better transform here, to ensure name is valid to the
	// CloudFormation name spec
	return strings.Replace(surveyAnswers.Name, " ", "", -1), nil
}

// surveyForEnvironments prompts the user to find out what environments this
// project uses
func surveyForEnvironments() (manifestEnvironments map[string]Environment, err error) {
	// Survey for which environments are used in this project
	environments := []string{}
	manifestEnvironments = map[string]Environment{}

	prompt := &survey.MultiSelect{
		Message: "Which environments does this project deploy to:",
		Help:    "you can add more later",
		Options: []string{"production", "staging", "development"},
	}
	// Prompts the user
	err = survey.AskOne(prompt, &environments, nil)
	if err != nil {
		return manifestEnvironments, err
	}
	for _, env := range environments {
		accountId, err := surveyForAccountId(env)
		if err != nil {
			return manifestEnvironments, err
		}
		manifestEnvironments[env] = Environment{
			AccountIDs: []string{accountId},
			Parameters: map[string]string{"Environment": env},
		}
	}

	return manifestEnvironments, err
}

// surveyForAccountId prompts the user to find out what accounts each environment uses
func surveyForAccountId(environment string) (accountId string, err error) {
	prompt := &survey.Input{
		Message: fmt.Sprintf("What is the Account ID for %s:", environment),
		Help:    "This is a whitelist of accounts, these stacks and parameters can be deployed to. This can prevent unintentional deployment.",
	}
	// Prompts the user
	err = survey.AskOne(prompt, &accountId, nil)
	if err != nil {
		return accountId, err
	}

	return accountId, err
}
