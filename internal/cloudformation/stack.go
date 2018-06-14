package cloudformation

import (
	"fmt"

	"github.com/KablamoOSS/kombustion/internal/manifest"
)

// GetStackName
// must satisfy regular expression pattern: [a-zA-Z][-a-zA-Z0-9]*
func GetStackName(manifestFile *manifest.Manifest, fileName, environment, stackNameFlag string) string {
	stackName := ""
	if stackNameFlag != "" {
		stackName = stackNameFlag
	} else {
		// TODO: remove the ext
		fileNameCleaned := fileName

		stackName = fmt.Sprintf("%s-%s-%s", manifestFile.Name, fileNameCleaned, environment)
	}

	// TODO: parse stackName to ensure it meets cfn regex requirements (strip bad chars)
	return stackName
}
