package tasks

import (
	"strings"
)

// Converts a cli StringSlice in "Key=Value" format to a map
func cliSliceMap(flags []string) map[string]string {
	flagMap := make(map[string]string)
	for _, flag := range flags {
		parts := strings.SplitN(flag, "=", 2)
		if len(parts) > 1 {
			flagMap[parts[0]] = parts[1]
		}
	}
	return flagMap
}
