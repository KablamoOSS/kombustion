package manifest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type initialiseTestPrompter struct {
	Name              string
	NameError         error
	Environments      []string
	EnvironmentsError error
	AccountIDs        map[string]string
	AccountIDsError   error
}

func (itp *initialiseTestPrompter) name() (string, error) {
	if itp.NameError != nil {
		return "", itp.NameError
	}
	return itp.Name, nil
}

func (itp *initialiseTestPrompter) environments() ([]string, error) {
	if itp.EnvironmentsError != nil {
		return []string{}, itp.EnvironmentsError
	}
	return itp.Environments, nil
}

func (itp *initialiseTestPrompter) accountID(env string) (string, error) {
	if itp.AccountIDsError != nil {
		return "", itp.AccountIDsError
	}
	return itp.AccountIDs[env], nil
}

func TestSurveyForInitialManifestHappyPath(t *testing.T) {
	testPrompt := &initialiseTestPrompter{
		Name:         "Kombustion",
		Environments: []string{"ci"},
		AccountIDs: map[string]string{
			"ci": "12345",
		},
	}

	name, environments, err := surveyForInitialManifest(testPrompt)
	assert.Nil(t, err)
	assert.Equal(t, name, "Kombustion")
	assert.Equal(t, len(environments), 1)
	assert.Equal(t, len(environments["ci"].AccountIDs), 1)
	assert.Equal(t, environments["ci"].AccountIDs[0], "12345")
}

func TestSurveyForInitialManifestError(t *testing.T) {
	testPrompt := &initialiseTestPrompter{
		Name:         "Komb",
		NameError:    fmt.Errorf("aborted"),
		Environments: []string{"N/A"},
		AccountIDs:   map[string]string{},
	}

	name, environments, err := surveyForInitialManifest(testPrompt)
	assert.Equal(t, err.Error(), "aborted")
	assert.Equal(t, name, "")
	assert.Equal(t, len(environments), 0)
}
