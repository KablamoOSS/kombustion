package manifest

import (
	"fmt"
	"testing"

	"github.com/KablamoOSS/kombustion/internal/coretest"
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

	manifest, err := surveyForInitialManifest(testPrompt)
	assert.Nil(t, err)
	assert.Equal(t, manifest.Name, "Kombustion")
	assert.Equal(t, len(manifest.Environments), 1)
	assert.Equal(t, len(manifest.Environments["ci"].AccountIDs), 1)
	assert.Equal(t, manifest.Environments["ci"].AccountIDs[0], "12345")
}

func TestSurveyForInitialManifestError(t *testing.T) {
	testPrompt := &initialiseTestPrompter{
		Name:         "Komb...",
		NameError:    fmt.Errorf("aborted"),
		Environments: []string{"N/A"},
		AccountIDs:   map[string]string{},
	}

	manifest, err := surveyForInitialManifest(testPrompt)
	assert.Equal(t, err.Error(), "aborted")
	assert.Nil(t, manifest)
}

func TestSurveyForInitialManifestEmptyAccountID(t *testing.T) {
	testPrompt := &initialiseTestPrompter{
		Name:         "Kombustion",
		Environments: []string{"ci"},
		AccountIDs: map[string]string{
			"ci": "",
		},
	}

	manifest, err := surveyForInitialManifest(testPrompt)
	assert.Nil(t, err)
	assert.Equal(t, manifest.Name, "Kombustion")
	assert.Equal(t, len(manifest.Environments), 1)
	assert.Equal(t, len(manifest.Environments["ci"].AccountIDs), 0)
}

func TestInitialiseManifest(t *testing.T) {
	testPrompt := &initialiseTestPrompter{
		Name:         "Kombustion",
		Environments: []string{"ci"},
		AccountIDs: map[string]string{
			"ci": "12345",
		},
	}

	objectStore := coretest.NewMockObjectStore()

	err := initialiseNewManifest(objectStore, testPrompt)
	assert.Nil(t, err)

	data, err := objectStore.Get("kombustion.yaml")
	assert.Nil(t, err)

	assert.Contains(t, string(data), "Name: Kombustion\n")
	assert.Contains(t, string(data), "Environments:\n  ci")
	assert.Contains(t, string(data), "GenerateDefaultOutputs: false\n")
}
