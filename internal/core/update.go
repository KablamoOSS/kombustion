package core

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/KablamoOSS/go-cli-printer"
	"github.com/google/go-github/github"
	"gopkg.in/AlecAivazis/survey.v1"
)

var githubClient *github.Client

func init() {
	githubClient = github.NewClient(nil)
}

// Update kombustion with the latest release on Github
func Update(currentVersion string, noPrompt bool) {
	printer.Step("Updating kombustion")
	printer.SubStep(
		fmt.Sprintf("Current Version: %s", currentVersion),
		1,
		false,
		true,
	)
	printer.Progress("Getting latest version")

	// 1. Check if there is an update
	latestRelease, downloadURL, foundRelease := checkForUpdate(currentVersion)
	// 2. If update, download update
	if foundRelease {
		printer.Stop()

		if *latestRelease.TagName == currentVersion {
			printer.SubStep(
				fmt.Sprintf("You have the current version."),
				1,
				true,
				true,
			)
			return
		}

		printer.SubStep(
			fmt.Sprintf("Found release %s", *latestRelease.TagName),
			1,
			false,
			true,
		)

		var confirm bool

		if noPrompt == false {
			confirm = surveyToConfirmUpdate(*latestRelease.TagName)
		} else {
			confirm = true
		}

		if confirm {
			printer.Progress(fmt.Sprintf("Downloading %s", *latestRelease.TagName))

			// TODO: Codesign the update, and check it here

			// 3. Replace current executable with new file
			downloadRelease(downloadURL)

			printer.SubStep(
				fmt.Sprintf("Update successful."),
				1,
				true,
				true,
			)
		} else {
			printer.SubStep(
				fmt.Sprintf("No update performed."),
				1,
				true,
				true,
			)
			return
		}
	} else {
		printer.SubStep(
			fmt.Sprintf("No updates found."),
			1,
			true,
			true,
		)
	}
}

func checkForUpdate(
	currentVersion string,
) (
	latestRelease *github.RepositoryRelease,
	downloadURL string,
	foundRelease bool,
) {
	latestRelease, err := GetLatestRelease(
		githubClient,
		"KablamoOSS",
		"kombustion",
	)
	if err != nil {
		printer.Fatal(
			err,
			"This may have failed due to network connectivity. Try updating again, but it fails download the latest version directly.",
			"https://www.kombustion.io/docs/downloads/",
		)
	}

	downloadURL, foundRelease = getReleaseForCurrentArchitecture(latestRelease)
	return latestRelease, downloadURL, foundRelease
}

func getReleaseForCurrentArchitecture(
	latestRelease *github.RepositoryRelease,
) (
	downloadURL string,
	foundRelease bool,
) {
	currentOS := runtime.GOOS
	currentArch := runtime.GOARCH

	if latestRelease.Assets != nil {
		for _, release := range latestRelease.Assets {
			operatingSystem, architecture, valid := GetOSArchFromFilename(
				"kombustion",
				*release.Name,
			)

			if valid == true &&
				operatingSystem == currentOS &&
				architecture == currentArch {
				downloadURL = *release.BrowserDownloadURL
				foundRelease = true
			}
		}
	}
	return downloadURL, foundRelease
}

func downloadRelease(downloadURL string) {

	urlSplit := strings.Split(downloadURL, "/")
	fileName := urlSplit[len(urlSplit)-1]

	response, err := http.Get(downloadURL)
	if err != nil {

	}
	defer response.Body.Close()

	printer.SubStep(
		fmt.Sprintf("Downloaded %s", downloadURL),
		1,
		false,
		true,
	)
	printer.Progress("Extracting")

	file, _ := ioutil.TempFile("", fileName)
	if err != nil {

	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {

	}

	file.Chmod(0755)

	extractedFolderPath, err := extractRelease(downloadURL, file.Name())

	newBinary := fmt.Sprintf("%s/kombustion", extractedFolderPath)

	// Executable read only
	os.Chmod(newBinary, 1400)

	currentExecutable, _ := os.Executable()
	// We expected the binary to be called `kombustion`
	os.Rename(newBinary, currentExecutable)

	printer.SubStep(
		fmt.Sprintf("Updated %s", currentExecutable),
		1,
		false,
		true,
	)
}

func extractRelease(url, fileName string) (string, error) {
	destination, _ := ioutil.TempDir("", "kombustion-update")

	extracter := GetExtracter(url)
	if extracter == nil {
		return "", fmt.Errorf(fmt.Sprintf("Unable to extract: %s", fileName))
	}
	err := extracter.Open(fileName, destination)
	if err != nil {
		printer.Fatal(
			err,
			"An error occured installing the new release. Try again, or download the latest release directly.",
			"https://www.kombustion.io/docs/downloads/",
		)
	}

	return destination, err
}

func surveyToConfirmUpdate(version string) bool {
	var update bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf(" ├─ Update to version %s?", version),
	}
	survey.AskOne(prompt, &update, nil)
	return update
}
