package tasks

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/KablamoOSS/kombustion/internal/coretest"
)

func TestGenerateTemplates(t *testing.T) {
	objectStore := coretest.NewMockObjectStore()
	objectStore.Put([]byte(sampleKombLock), "kombustion.lock")

	tests := []struct {
		templateFile      string
		manifestFile      string
		generatedTemplate string
	}{
		// To add new test cases, compile kombution and generate a template from the source template
		// with a command similar to `kombustion --manifest-file S3AccessLogs.manifest.yaml generate S3AccessLogs.yaml`
		{
			templateFile:      "testdata/generate/S3AccessLogs.yaml",
			manifestFile:      "testdata/generate/S3AccessLogs.manifest.yaml",
			generatedTemplate: "testdata/generate/S3AccessLogs.generated.yaml",
		},
	}

	for i, test := range tests {
		assert := assert.New(t)

		objectStore.PutFile(test.templateFile, "template.yaml")
		objectStore.PutFile(test.manifestFile, "kombustion.yaml")
		objectStore.PutFile(test.generatedTemplate, "expected.yaml")

		generate(
			objectStore,         // objectStore
			"template.yaml",     // templatePath
			map[string]string{}, // cliParams
			"",                  // paramsPath
			"",                  // devPluginPath
			"compiled",          // outputDirectory
			false,               // ouputParameters
			"ci",                // envName
			false,               // generateDefaultOutputs
			"kombustion.yaml",   // manifest location
		)

		output, _ := objectStore.Get("compiled", "template.yaml")

		expectedOutput, _ := objectStore.Get("expected.yaml")
		assert.Equal(string(expectedOutput), string(output), fmt.Sprintf("Failed generating template: Test #%d", i))
	}
}
