package plugins

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	manifestType "github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
	"github.com/google/go-github/github"
)

var githubClient *github.Client

func init() {
	githubClient = github.NewClient(nil)
}

// AddPluginsToManifest - Add all new plugin to the manifest
// update it if it's already there
// then write the manifest to disk
func AddPluginsToManifest(manifest *manifestType.Manifest, pluginLocations []string) (*manifestType.Manifest, error) {
	printer.Progress("Kombusting")

	// Get the lockFile
	lockFile, err := lock.FindAndLoadLock()
	if err != nil {
		printer.Error(err, config.ErrorHelpInfo, "")
		return manifest, err
	}

	// Add all the plugins to the manifest and lockfile
	manifest, lockFile, err = addPluginsToManifestAndLock(manifest, lockFile, pluginLocations)
	if err != nil {
		printer.Error(err, config.ErrorHelpInfo, "")
		return manifest, err
	}

	printer.Progress("Updating manifest")
	err = manifestType.WriteManifestToDisk(manifest)
	if err != nil {
		printer.Error(err, config.ErrorHelpInfo, "")
		return manifest, err
	}

	err = lock.UpdateLock(manifest, lockFile)
	if err != nil {
		printer.Error(err, fmt.Sprintf("%s\n%s", "There was an error updating lockfile. Try your previous command again, and if that fails you may need to remove kombustion.lock entirely.", config.ErrorHelpInfo), "")
		return manifest, err
	}

	return manifest, nil
}

func addPluginsToManifestAndLock(
	manifest *manifestType.Manifest,
	lockFile *lock.Lock,
	pluginLocations []string,
) (
	*manifestType.Manifest,
	*lock.Lock,
	error,
) {
	for _, pluginLocation := range pluginLocations {
		plugin, pluginLock, err := constructGithubPlugin(manifest, pluginLocation)
		printer.SubStep(fmt.Sprintf("Adding plugin: %s", plugin.Name), 2, true, false)
		if err != nil {
			return manifest, lockFile, err
		}

		if manifest.Plugins == nil {
			manifest.Plugins = make(map[string]manifestType.Plugin)
		}
		if lockFile.Plugins == nil {
			lockFile.Plugins = make(map[string]lock.Plugin)
		}

		manifest.Plugins[fmt.Sprintf("%s@%s", plugin.Name, plugin.Version)] = plugin
		lockFile.Plugins[fmt.Sprintf("%s@%s", plugin.Name, plugin.Version)] = pluginLock
	}
	return manifest, lockFile, nil
}

// constructGithubPlugin - Create a plugin based on a github url
func constructGithubPlugin(
	manifest *manifestType.Manifest, pluginURI string,
) (
	plugin manifestType.Plugin,
	pluginLock lock.Plugin,
	err error,
) {
	var pluginURL *url.URL
	// Parse the plugin url
	pluginURL, err = url.Parse(pluginURI)
	if err != nil {
		return plugin, pluginLock, err
	}

	path := strings.Split(pluginURL.Path, "/")
	if path[0] != "github.com" {
		return plugin, pluginLock, fmt.Errorf("Plugin must be start with github.com")
	}
	githubOrg := path[1]
	githubProject := path[2]

	plugin.Name = strings.Join([]string{
		pluginURL.Host,
		pluginURL.Path,
	},
		"",
	)
	pluginLock.Name = plugin.Name

	latestRelease, latestReleaseErr := getLatestRelease(githubOrg, githubProject)
	// TODO: handle no release
	// TODO: handle release with no files
	if err != nil {
		return plugin, pluginLock, latestReleaseErr
	}

	if pluginLock.Resolved == nil {
		pluginLock.Resolved = make(map[string]lock.PluginResolution)
	}

	plugin.Version = *latestRelease.TagName
	pluginLock.Version = plugin.Version

	// Loop through the assets we found and try to create a resolution lock for them
	if latestRelease.Assets != nil {
		for _, release := range latestRelease.Assets {
			operatingSystem, architecture, valid := getOSArchFromFilename(githubProject, *release.Name)

			// If the file is a valid plugin file create a resolution lock
			if valid {
				pluginLock.Resolved[strings.Join([]string{operatingSystem, architecture}, "-")] = lock.PluginResolution{
					URL:             *release.BrowserDownloadURL,
					OperatingSystem: operatingSystem,
					Architecture:    architecture,
					PathOnDisk:      "",
					Hash:            "",
					ArchiveHash:     "",
					ArchiveName:     *release.Name,
				}
			}
		}
	}

	return plugin, pluginLock, err
}

// getLatestRelease - Return the latest release of the repository
func getLatestRelease(
	githubOrg string,
	githubProject string,
) (
	latestRelease *github.RepositoryRelease,
	err error,
) {

	// Get latest release
	latestRelease, _, err = githubClient.Repositories.GetLatestRelease(
		context.Background(),
		githubOrg,
		githubProject,
	)
	if err != nil {
		return latestRelease, err
	}

	printer.SubStep(
		fmt.Sprintf("Found release %s for %s/%s", *latestRelease.TagName, githubOrg, githubProject),
		1,
		true,
		false,
	)

	return latestRelease, nil
}

// getOSArchFromFilename - Extract the os and arch from the file name
// Expecting the filename to be of the format `{pluginName}-{os}-{arch}.{tgz|zip}`
func getOSArchFromFilename(pluginName string, fileName string) (operatingSystem string, architecture string, valid bool) {
	// The filename must start with the plugin name
	if strings.HasPrefix(fileName, strings.Join([]string{pluginName, "-"}, "")) == false {
		// Err, this file is not a plugin
		return "", "", false
	}

	// We expect .tgz or .zip as the extension so remove both
	fileNameWithoutExtension := strings.Replace(strings.Replace(fileName, ".tgz", "", 1), ".zip", "", 1)

	// first remove the plugin name from the filename to get the osArch string
	osArch := strings.Replace(
		fileNameWithoutExtension,
		strings.Join([]string{pluginName, "-"}, ""),
		"",
		1,
	)

	osArchSplit := strings.Split(osArch, "-")

	operatingSystem = osArchSplit[0]
	// The last item should be the arch, for cases of darwin-10.6-amd64
	architecture = osArchSplit[len(osArchSplit)-1]

	// Check that the os and arch match something go can build
	if checkValidOS(operatingSystem) && checkValidArch(architecture) {
		return operatingSystem, architecture, true
	}

	return "", "", false
}

//  checkValidOS - Check if the input is a valid Go OS target
func checkValidOS(input string) (valid bool) {
	validOS := []string{
		"darwin",
		"freebsd",
		"linux",
		// Note windows doesn't actually have plugin support yet
		"windows",
	}

	for _, OS := range validOS {
		if OS == input {
			valid = true
			return valid
		}
	}
	return valid
}

//  checkValidArch - Check if the input is a valid Go arch target
func checkValidArch(input string) (valid bool) {
	validArch := []string{
		"386",
		"amd64",
		"arm64",
	}

	for _, arch := range validArch {
		if arch == input {
			valid = true
			return valid
		}
	}
	return valid
}
