package cloudformation

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/KablamoOSS/kombustion/internal/manifest"
)

// GetStackName cleaned to fix AWS requirements
func GetStackName(manifestFile *manifest.Manifest, fileName, environment, stackNameFlag string) string {
	stackName := ""
	if stackNameFlag != "" {
		stackName = stackNameFlag
	} else {
		// Remove the folder and extension, leaving only the filename
		fileNameCleaned := strings.Replace(path.Base(fileName), path.Ext(fileName), "", 1)

		stackName = cleanStackName(fmt.Sprintf(
			"%s-%s-%s",
			cleanStackName(manifestFile.Name),
			cleanStackName(fileNameCleaned),
			cleanStackName(environment),
		),
		)
	}

	return stackName
}

// Proceses a string to ensure it matches
// must satisfy regular expression pattern: [a-zA-Z][-a-zA-Z0-9]*
func cleanStackName(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9\\-]+")

	return reg.ReplaceAllString(strings.Replace(input, "/", "-", 0), "")
}
