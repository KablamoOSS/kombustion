package plugins

import (
	"fmt"

	"os"
	"os/user"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
)

func findPluginInCache(plugin lock.Plugin, resolved lock.PluginResolution) (found bool, cacheFile string, err error) {
	printer.Progress(
		fmt.Sprintf(
			"Checking download cache for %s@%s[%s/%s]",
			plugin.Name,
			plugin.Version,
			resolved.OperatingSystem,
			resolved.Architecture,
		))
	cacheFile = fmt.Sprintf("%s/%s", getDownloadDir(plugin.Name, plugin.Version), resolved.ArchiveName)
	if _, err := os.Stat(cacheFile); err == nil {
		// cacheFile exists
		cacheFileHash, err := getHashOfFile(cacheFile)
		if err != nil {
			return found, cacheFile, err
		}
		if cacheFileHash == resolved.ArchiveHash {
			// We have a matching downloaded file
			found = true
		}
	}
	return found, cacheFile, err
}

// GetCacheDir - Get the users cache directory
// Make it if it doesn't exist
func GetCacheDir() string {
	usr, err := user.Current()
	if err != nil {
		// TODO: Make this error more helpful
		printer.Fatal(
			err,
			"",
			"",
		)
	}
	plugindir := fmt.Sprintf("%s/.kombustion/cache/plugins", usr.HomeDir)
	os.MkdirAll(plugindir, 0744)

	return plugindir
}

func installFromCache(cacheFile string, plugin lock.Plugin, resolved lock.PluginResolution) (couldInstall bool, installErrors []error) {
	extractedFileName, extractErr := extractPlugin(
		plugin.Name,
		resolved.OperatingSystem,
		resolved.Architecture,
		plugin.Version,
		cacheFile,
	)
	if extractErr != nil {
		installErrors = append(installErrors, extractErr)
	}
	hash, hashErr := getHashOfFile(extractedFileName)
	if hashErr != nil {
		installErrors = append(installErrors, hashErr)
	}
	if hash == resolved.Hash {
		couldInstall = true
	} else {
		// If the hash fails, remove th extracted file
		os.Remove(extractedFileName)
	}
	return couldInstall, installErrors
}
