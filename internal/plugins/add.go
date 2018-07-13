package plugins

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/config"
	"github.com/KablamoOSS/kombustion/internal/core"
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
	lockFile := lock.FindAndLoadLock()
	var err error

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

	latestRelease, latestReleaseErr := core.GetLatestRelease(githubClient, githubOrg, githubProject)
	if latestReleaseErr != nil {
		// TODO: Make this error more helpful
		printer.Fatal(
			latestReleaseErr,
			"",
			"",
		)
	}

	if latestRelease != nil {
		printer.SubStep(
			fmt.Sprintf("Found release %s for %s/%s", *latestRelease.TagName, githubOrg, githubProject),
			1,
			true,
			false,
		)

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
				operatingSystem, architecture, valid := core.GetOSArchFromFilename(githubProject, *release.Name)

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
	}
	return plugin, pluginLock, err
}
