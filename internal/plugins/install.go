package plugins

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/core"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
)

// InstallPlugins - Get the lock file and then call installPluginsWithLock
func InstallPlugins() error {
	printer.Step("Installing plugins")
	printer.Progress("Kombusting")

	lockFile := lock.FindAndLoadLock()

	updatedLockFile, installErrors := installPluginsWithLock(lockFile)
	if len(installErrors) > 0 {
		for _, err := range installErrors {
			printer.Error(err, "", "")
		}

		// TODO: This error message could be more helpful
		printer.Fatal(
			fmt.Errorf("failed installing plugins"),
			"Error installing plugins",
			"",
		)
	}
	err := lock.WriteLockToDisk(updatedLockFile)
	if err != nil {
		printer.Fatal(err, "Error installing plugins", "")
		return err
	}
	return nil
}

// installPluginsWithLock - Using a provided lockFile install plugins
func installPluginsWithLock(lockFile *lock.Lock) (updatedLockFile *lock.Lock, installErrors []error) {
	updatedLockFile = *&lockFile

	ensureLocalPluginDir()

	for i, plugin := range lockFile.Plugins {
		updatedPlugin, errs := installPlugin(plugin)
		if errs != nil {
			installErrors = append(installErrors, errs...)
		}
		updatedLockFile.Plugins[i] = updatedPlugin
	}

	if len(installErrors) > 0 {
		printer.Error(installErrors[0], "Failed to install plugins", "")
		return updatedLockFile, installErrors
	}
	printer.Finish("Installed plugins")
	return updatedLockFile, installErrors
}

// installPlugin - Install an individual plugin
func installPlugin(plugin lock.Plugin) (updatedPlugin lock.Plugin, installErrors []error) {
	printer.SubStep(fmt.Sprintf("Installing %s", plugin.Name), 1, false, false)

	updatedPlugin = plugin
	for i, resolved := range plugin.Resolved {
		pluginIsInstalled := checkPluginIsAlreadyInstalled(plugin, resolved)
		if pluginIsInstalled {
			printer.SubStep(fmt.Sprintf(
				"Already installed %s@%s[%s/%s]",
				plugin.Name,
				plugin.Version,
				resolved.OperatingSystem,
				resolved.Architecture,
			),
				2,
				true,
				false,
			)
		} else {
			var couldInstallFromCache bool
			updatedResolved := resolved

			// Check the local cache for a file
			foundInCache, cacheFile, err := findPluginInCache(plugin, resolved)
			if err != nil {
				printer.Fatal(
					err,
					"Try again, if this error persists try clearing the plugin cache.",
					"https://www.kombustion.io/concepts/plugins/#cache",
				)
				installErrors = append(installErrors, err)
			}

			if foundInCache {
				printer.SubStep(fmt.Sprintf(
					"Found cache for %s@%s[%s/%s]",
					plugin.Name,
					plugin.Version,
					resolved.OperatingSystem,
					resolved.Architecture,
				),
					2,
					false,
					false,
				)

				var cacheErrors []error
				couldInstallFromCache, cacheErrors = installFromCache(cacheFile, plugin, resolved)
				if len(cacheErrors) != 0 {
					installErrors = append(installErrors, cacheErrors...)
				}
			}
			if couldInstallFromCache == false {
				var downloadErrors []error
				updatedResolved, downloadErrors = downloadPlugin(plugin, resolved)
				if len(downloadErrors) != 0 {
					installErrors = append(installErrors, downloadErrors...)
				}
			}
			updatedPlugin.Resolved[i] = updatedResolved

		}
	}

	printer.SubStep(fmt.Sprintf("Installed %s", plugin.Name), 2, true, false)
	return updatedPlugin, installErrors
}

func checkPluginIsAlreadyInstalled(plugin lock.Plugin, resolved lock.PluginResolution) (pluginIsInstalled bool) {
	// If there's no hash yet, it means the file isn't installed
	if resolved.Hash == "" {
		return false
	}

	hash, _ := getHashOfFile(resolved.PathOnDisk)
	if hash == resolved.Hash {
		pluginIsInstalled = true
	}
	return pluginIsInstalled
}

// downloadPlugin - Download a single plugin
func downloadPlugin(plugin lock.Plugin, resolved lock.PluginResolution) (updatedResolved lock.PluginResolution, downloadErrors []error) {
	printer.Progress(
		fmt.Sprintf(
			"Downloading %s@%s[%s/%s]",
			plugin.Name,
			plugin.Version,
			resolved.OperatingSystem,
			resolved.Architecture,
		),
	)

	// TODO: check the URL is valid

	urlSplit := strings.Split(resolved.URL, "/")
	fileName := urlSplit[len(urlSplit)-1]
	filePath := fmt.Sprintf("%s/%s", getDownloadDir(plugin.Name, plugin.Version), fileName)

	resolved.ArchiveName = fileName

	output, err := os.Create(filePath)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	defer output.Close()

	response, err := http.Get(resolved.URL)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	defer response.Body.Close()

	// Save the downloaded file
	_, err = io.Copy(output, response.Body)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}

	resolved.ArchiveHash, err = getHashOfFile(filePath)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}
	extractedFileName, extractErr := extractPlugin(
		plugin.Name,
		resolved.OperatingSystem,
		resolved.Architecture,
		plugin.Version,
		filePath,
	)
	if extractErr != nil {
		downloadErrors = append(downloadErrors, extractErr)
	}

	resolved.PathOnDisk = extractedFileName

	resolved.Hash, err = getHashOfFile(resolved.PathOnDisk)
	if err != nil {
		downloadErrors = append(downloadErrors, err)
	}

	printer.SubStep(
		fmt.Sprintf(
			"Downloaded %s@%s[%s/%s]",
			plugin.Name,
			plugin.Version,
			resolved.OperatingSystem,
			resolved.Architecture,
		),
		2,
		false,
		false,
	)

	updatedResolved = resolved
	return updatedResolved, downloadErrors
}

// Extract a plugin downloaded to the local cache, into the local plugin dir
func extractPlugin(pluginName string, operatingSystem string, architecture string, version string, fileName string) (extractedFilePath string, err error) {
	destination := getLocalPluginDir(pluginName, operatingSystem, architecture, version)
	extracter := core.GetExtracter(fileName)
	if extracter == nil {
		return extractedFilePath, fmt.Errorf(fmt.Sprintf("Unable to extract: %s", fileName))
	}
	err = extracter.Open(fileName, destination)
	var found bool
	var extractedFileName string
	extractedFileName, found, err = findExtractedPluginFile(destination)
	if found == false {
		return extractedFilePath, fmt.Errorf("Invalid plugin archive: %s", pluginName)
	}
	extractedFilePath = fmt.Sprintf("%s/%s", destination, extractedFileName)
	return extractedFilePath, err
}

func findExtractedPluginFile(directory string) (fileName string, found bool, err error) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return fileName, found, err
	}

	for _, file := range files {
		fileName = file.Name()
		if strings.HasSuffix(fileName, ".so") {
			found = true
			return fileName, found, err
		}
	}
	return fileName, found, err
}

// getPluginDir - Get the plugin directory for this project,
// and make it if it doesn't exist
func getLocalPluginDir(pluginName string, operatingSystem string, architecture string, version string) string {
	pluginDir := fmt.Sprintf(".kombustion/plugins/%s/%s/%s/%s", pluginName, version, operatingSystem, architecture)
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}

func ensureLocalPluginDir() string {
	pluginDir := fmt.Sprintf(".kombustion/plugins/")
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}

// Get the download cache directory
func getDownloadDir(pluginName string, version string) string {
	cacheDir := GetCacheDir()
	pluginDir := fmt.Sprintf("%s/%s/%s", cacheDir, pluginName, version)
	os.MkdirAll(pluginDir, 0744)
	return pluginDir
}
