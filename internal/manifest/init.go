package manifest

import (
	"fmt"
	"strings"

	"github.com/KablamoOSS/kombustion/internal/core"
	printer "github.com/KablamoOSS/go-cli-printer"
	"gopkg.in/AlecAivazis/survey.v1"
)

type initialisePrompter interface {
	name() (string, error)
	environments() ([]string, error)
	accountID(string) (string, error)
}

// InitaliseNewManifest creates a new manifest with a survey
func InitaliseNewManifest(objectStore core.ObjectStore) error {
	// Load the manifest file from this directory
	if CheckManifestExists(objectStore) {
		printer.Fatal(
			fmt.Errorf("Sorry we can't create a new kombustion.yaml, one already exists."),
			"If you want to re-initialise your kombustion.yaml file, first remove it.",
			"https://www.kombustion.io/api/manifest/",
		)
	}

	// Survey the user for required info
	manifest, err := surveyForInitialManifest(&surveyPrompt{})
	if err != nil {
		return err
	}

	err = WriteManifestToDisk(manifest)
	if err != nil {
		return err
	}
	return nil
}

// surveyForInitialManifest - Prompt the user to fill out the required fields,
// but check if we have a flag for them
func surveyForInitialManifest(prompter initialisePrompter) (*Manifest, error) {
	manifest := &Manifest{
		Environments: map[string]Environment{},
	}

	name, err := prompter.name()
	if err != nil {
		return nil, err
	}
	manifest.Name = name

	environmentNames, err := prompter.environments()
	if err != nil {
		return nil, err
	}

	for _, env := range environmentNames {
		accountId, err := prompter.accountID(env)
		if err != nil {
			return nil, err
		}

		// Adding an empty account ID into kombustion.yaml would mean that no
		// account would be valid, and kombustion would refuse to perform any
		// (AWS) operation. Instead, if the user doesn't enter an ID, assume
		// they don't want a whitelist.
		accountIds := []string{}
		if accountId != "" {
			accountIds = append(accountIds, accountId)
		}

		manifest.Environments[env] = Environment{
			AccountIDs: accountIds,
			Parameters: map[string]string{"Environment": env},
		}
	}

	return manifest, nil
}

// Implements the initialisePrompter interface, so we can plug another
// implementation in for testing
type surveyPrompt struct{}

// surveyForName - Prompt for the name of the project
func (sp *surveyPrompt) name() (string, error) {
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
func (sp *surveyPrompt) environments() ([]string, error) {
	// Survey for which environments are used in this project
	environments := []string{}

	prompt := &survey.MultiSelect{
		Message: "Which environments does this project deploy to:",
		Help:    "you can add more later",
		Options: []string{"production", "staging", "development"},
	}
	// Prompts the user
	err := survey.AskOne(prompt, &environments, nil)
	if err != nil {
		return environments, err
	}

	return environments, err
}

// surveyForAccountId prompts the user to find out what accounts each environment uses
func (sp *surveyPrompt) accountID(environment string) (accountId string, err error) {
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
