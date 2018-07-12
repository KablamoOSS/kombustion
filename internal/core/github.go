package core

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"github.com/mholt/archiver"
)

// GetLatestRelease - Return the latest release of the repository
func GetLatestRelease(
	githubClient *github.Client,
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

	return latestRelease, nil
}

// GetOSArchFromFilename - Extract the os and arch from the file name
// Expecting the filename to be of the format `{name}-{os}-{arch}.{tgz|zip}`
func GetOSArchFromFilename(name string, fileName string) (operatingSystem string, architecture string, valid bool) {
	// The filename must start with the plugin name
	if strings.HasPrefix(fileName, strings.Join([]string{name, "-"}, "")) == false {
		// Err, this file is not a plugin
		return "", "", false
	}

	// We expect .tgz or .zip as the extension so remove both
	fileNameWithoutExtension := strings.Replace(strings.Replace(fileName, ".tgz", "", 1), ".zip", "", 1)

	// first remove the plugin name from the filename to get the osArch string
	osArch := strings.Replace(
		fileNameWithoutExtension,
		strings.Join([]string{name, "-"}, ""),
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

// GetExtracter The PR hasn't been released into a new version, so we're putting it here
// https://github.com/mholt/archiver/pull/45/files
func GetExtracter(fpath string) archiver.Archiver {
	for _, format := range archiver.SupportedFormats {
		if format.Match(fpath) {
			return format
		}
	}
	return nil
}
