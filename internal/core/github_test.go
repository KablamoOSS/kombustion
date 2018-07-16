package core

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/github"
	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setupMockGithubClient() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	githubClient = github.NewClient(nil)
	githubClient.BaseURL, _ = url.Parse(fmt.Sprintf("%s/", server.URL))
}

func teardownMockGithubClient() {
	server.Close()
}

func TestGetLatestRelease(t *testing.T) {
	type input struct {
		githubOrg     string
		githubProject string
	}
	type output struct {
		version string
		err     error
	}
	tests := []struct {
		input     input
		output    output
		throws    bool
		httpFuncs map[string]func(http.ResponseWriter, *http.Request)
	}{
		{
			throws: false,
			input: input{
				githubOrg:     "KablamoOSS",
				githubProject: "kombustion-plugin-test",
			},
			httpFuncs: map[string]func(http.ResponseWriter, *http.Request){
				"/repos/KablamoOSS/kombustion-plugin-test/releases/latest": func(w http.ResponseWriter, r *http.Request) {
					data, err := ioutil.ReadFile("testdata/latestRelease.json")
					if err != nil {
						t.Fail()
					}

					fmt.Fprint(w, string(data))
				},
			},
			output: output{
				version: "0.1.0",
				err:     nil,
			},
		},
	}

	for i, test := range tests {
		setupMockGithubClient()
		defer teardownMockGithubClient()
		assert := assert.New(t)

		for url, handler := range test.httpFuncs {
			mux.HandleFunc(url, handler)
		}

		latestRelease, err := GetLatestRelease(
			githubClient,
			test.input.githubOrg,
			test.input.githubProject,
		)
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(
				*latestRelease.TagName,
				test.output.version,
				fmt.Sprintf("Test %d: %s", i, test.input),
			)
		}
	}
}

func TestGetOSArchFromFilename(t *testing.T) {
	type input struct {
		pluginName string
		fileName   string
	}
	type output struct {
		operatingSystem string
		architecture    string
		valid           bool
	}
	tests := []struct {
		input  input
		output output
	}{
		{
			input: input{
				pluginName: "kombustion-plugin-test",
				fileName:   "kombustion-plugin-test-darwin-10.6-amd64.tgz",
			},
			output: output{
				operatingSystem: "darwin",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-test",
				fileName:   "kombustion-plugin-test-linux-386.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "386",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-test",
				fileName:   "kombustion-plugin-test-linux-amd64.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-test",
				fileName:   "kombustion-plugin-test-linux-arm64.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "arm64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "invalid-plugin-name",
				fileName:   "kombustion-plugin-test-linux-arm64.tgz",
			},
			output: output{
				operatingSystem: "",
				architecture:    "",
				valid:           false,
			},
		},
		{
			input: input{
				pluginName: "",
				fileName:   "",
			},
			output: output{
				operatingSystem: "",
				architecture:    "",
				valid:           false,
			},
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		operatingSystem, architecture, valid := GetOSArchFromFilename(
			test.input.pluginName,
			test.input.fileName,
		)
		testOutput := output{
			operatingSystem: operatingSystem,
			architecture:    architecture,
			valid:           valid,
		}
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}

func TestCheckValidOS(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "darwin",
			output: true,
		},
		{
			input:  "freebsd",
			output: true,
		},

		{
			input:  "linux",
			output: true,
		},

		{
			input:  "fail",
			output: false,
		},

		{
			input:  "fail-123",
			output: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := checkValidOS(
			test.input,
		)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}

func TestCheckValidArch(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{
			input:  "386",
			output: true,
		},

		{
			input:  "380",
			output: false,
		},

		{
			input:  "amd64",
			output: true,
		},

		{
			input:  "arm64",
			output: true,
		},

		{
			input:  "fail",
			output: false,
		},

		{
			input:  "fail-123",
			output: false,
		},
	}

	for i, test := range tests {
		assert := assert.New(t)
		testOutput := checkValidArch(
			test.input,
		)
		assert.Equal(testOutput, test.output, fmt.Sprintf("Test %d: %s", i, test.input))
	}
}
