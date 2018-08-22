package lock

import (
	"runtime"

	"github.com/KablamoOSS/kombustion/internal/manifest"
)

func (plugin *Plugin) Match(m *manifest.Plugin) bool {
	return plugin.Name == m.Name && plugin.Version == m.Version
}

func (plugin *Plugin) Resolve(os, arch string) (string, bool) {
	for _, resolution := range plugin.Resolved {
		if resolution.OperatingSystem == os && resolution.Architecture == arch {
			return resolution.PathOnDisk, true
		}
	}
	return "", false
}

func (plugin *Plugin) ResolveForRuntime() (string, bool) {
	return plugin.Resolve(runtime.GOOS, runtime.GOARCH)
}
