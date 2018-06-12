package plugins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	manifestType "github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/KablamoOSS/kombustion/internal/plugins/lock"
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

func TestAddPluginsToManifestAndLock(t *testing.T) {
	// TODO: Add tests to cover the error cases
	type input struct {
		manifest        *manifestType.Manifest
		lockFile        *lock.Lock
		pluginLocations []string
	}
	type output struct {
		manifest *manifestType.Manifest
		lock     *lock.Lock
		err      error
	}
	tests := []struct {
		name      string
		input     input
		output    output
		throws    bool
		httpFuncs map[string]func(http.ResponseWriter, *http.Request)
	}{
		{
			name:   "Add github plugin",
			throws: false,
			input: input{
				manifest: &manifestType.Manifest{},
				lockFile: &lock.Lock{},
				pluginLocations: []string{
					"github.com/KablamoOSS/kombustion-plugin-serverless",
				},
			},
			httpFuncs: map[string]func(http.ResponseWriter, *http.Request){
				"/repos/KablamoOSS/kombustion-plugin-serverless/releases/latest": func(w http.ResponseWriter, r *http.Request) {
					data, err := ioutil.ReadFile("testdata/add/releases/latestRelease.json")
					if err != nil {
						t.Fail()
					}

					fmt.Fprint(w, string(data))
				},
			},
			output: output{
				manifest: &manifestType.Manifest{
					Name: "",
					Plugins: map[string]manifestType.Plugin{
						"github.com/KablamoOSS/kombustion-plugin-serverless@0.1.0": manifestType.Plugin{
							Name:    "github.com/KablamoOSS/kombustion-plugin-serverless",
							Version: "0.1.0",
							Alias:   "",
						},
					},
					Architectures:      []string(nil),
					Environments:       map[string]manifestType.Environment(nil),
					HideDefaultExports: false,
				},
				lock: &lock.Lock{
					Plugins: map[string]lock.Plugin{
						"github.com/KablamoOSS/kombustion-plugin-serverless@0.1.0": lock.Plugin{
							Name:    "github.com/KablamoOSS/kombustion-plugin-serverless",
							Version: "0.1.0",
							Resolved: map[string]lock.PluginResolution{
								"linux-386": lock.PluginResolution{
									URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-386.tgz",
									OperatingSystem: "linux",
									Architecture:    "386",
									PathOnDisk:      "",
									Hash:            "",
									ArchiveHash:     "",
									ArchiveName:     "kombustion-plugin-serverless-linux-386.tgz",
								},
								"linux-amd64": lock.PluginResolution{
									URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-amd64.tgz",
									OperatingSystem: "linux",
									Architecture:    "amd64",
									PathOnDisk:      "",
									Hash:            "",
									ArchiveHash:     "",
									ArchiveName:     "kombustion-plugin-serverless-linux-amd64.tgz",
								},
								"linux-arm64": lock.PluginResolution{
									URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-arm64.tgz",
									OperatingSystem: "linux",
									Architecture:    "arm64",
									PathOnDisk:      "",
									Hash:            "",
									ArchiveHash:     "",
									ArchiveName:     "kombustion-plugin-serverless-linux-arm64.tgz",
								},
								"darwin-amd64": lock.PluginResolution{
									URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
									OperatingSystem: "darwin",
									Architecture:    "amd64",
									PathOnDisk:      "",
									Hash:            "",
									ArchiveHash:     "",
									ArchiveName:     "kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
								},
							},
						},
					},
				},
				err: nil,
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

		manifest, lock, err := addPluginsToManifestAndLock(
			test.input.manifest,
			test.input.lockFile,
			test.input.pluginLocations,
		)
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(
				manifest,
				test.output.manifest,
				fmt.Sprintf("Test %d: %s", i, test.name),
			)
			assert.Equal(
				lock,
				test.output.lock,
				fmt.Sprintf("Test %d: %s", i, test.name),
			)
		}
	}
}

func TestConstructGithubPlugin(t *testing.T) {
	// TODO: Add tests to cover the error cases
	type input struct {
		manifest  *manifestType.Manifest
		pluginURI string
	}
	type output struct {
		plugin     manifestType.Plugin
		pluginLock lock.Plugin
		err        error
	}
	tests := []struct {
		name      string
		input     input
		output    output
		throws    bool
		httpFuncs map[string]func(http.ResponseWriter, *http.Request)
	}{
		{
			name:   "Add github plugin",
			throws: false,
			input: input{
				manifest:  &manifestType.Manifest{},
				pluginURI: "github.com/KablamoOSS/kombustion-plugin-serverless",
			},
			httpFuncs: map[string]func(http.ResponseWriter, *http.Request){
				"/repos/KablamoOSS/kombustion-plugin-serverless/releases/latest": func(w http.ResponseWriter, r *http.Request) {
					data, err := ioutil.ReadFile("testdata/add/releases/latestRelease.json")
					if err != nil {
						t.Fail()
					}

					fmt.Fprint(w, string(data))
				},
			},
			output: output{
				plugin: manifestType.Plugin{
					Name:    "github.com/KablamoOSS/kombustion-plugin-serverless",
					Version: "0.1.0", Alias: "",
				},
				pluginLock: lock.Plugin{
					Name:    "github.com/KablamoOSS/kombustion-plugin-serverless",
					Version: "0.1.0",
					Resolved: map[string]lock.PluginResolution{
						"darwin-amd64": lock.PluginResolution{
							URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
							OperatingSystem: "darwin",
							Architecture:    "amd64",
							PathOnDisk:      "",
							Hash:            "",
							ArchiveHash:     "",
							ArchiveName:     "kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
						},
						"linux-386": lock.PluginResolution{
							URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-386.tgz",
							OperatingSystem: "linux",
							Architecture:    "386",
							PathOnDisk:      "",
							Hash:            "",
							ArchiveHash:     "",
							ArchiveName:     "kombustion-plugin-serverless-linux-386.tgz",
						},
						"linux-amd64": lock.PluginResolution{
							URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-amd64.tgz",
							OperatingSystem: "linux",
							Architecture:    "amd64",
							PathOnDisk:      "",
							Hash:            "",
							ArchiveHash:     "",
							ArchiveName:     "kombustion-plugin-serverless-linux-amd64.tgz",
						},
						"linux-arm64": lock.PluginResolution{
							URL:             "https://github.com/KablamoOSS/kombustion-plugin-serverless/releases/download/0.1.0/kombustion-plugin-serverless-linux-arm64.tgz",
							OperatingSystem: "linux",
							Architecture:    "arm64",
							PathOnDisk:      "",
							Hash:            "",
							ArchiveHash:     "",
							ArchiveName:     "kombustion-plugin-serverless-linux-arm64.tgz",
						},
					},
				},
				err: nil,
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

		plugin, pluginLock, err := constructGithubPlugin(
			test.input.manifest,
			test.input.pluginURI,
		)
		if test.throws {
			assert.NotNil(err)
		} else {
			assert.Nil(err)
			assert.Equal(
				plugin,
				test.output.plugin,
				fmt.Sprintf("Test %d: %s", i, test.name),
			)
			assert.Equal(
				pluginLock,
				test.output.pluginLock,
				fmt.Sprintf("Test %d: %s", i, test.name),
			)
		}
	}
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
				githubProject: "kombustion-plugin-serverless",
			},
			httpFuncs: map[string]func(http.ResponseWriter, *http.Request){
				"/repos/KablamoOSS/kombustion-plugin-serverless/releases/latest": func(w http.ResponseWriter, r *http.Request) {
					data, err := ioutil.ReadFile("testdata/add/releases/latestRelease.json")
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

		latestRelease, err := getLatestRelease(
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
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-darwin-10.6-amd64.tgz",
			},
			output: output{
				operatingSystem: "darwin",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-386.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "386",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-amd64.tgz",
			},
			output: output{
				operatingSystem: "linux",
				architecture:    "amd64",
				valid:           true,
			},
		},
		{
			input: input{
				pluginName: "kombustion-plugin-serverless",
				fileName:   "kombustion-plugin-serverless-linux-arm64.tgz",
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
				fileName:   "kombustion-plugin-serverless-linux-arm64.tgz",
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
		operatingSystem, architecture, valid := getOSArchFromFilename(
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
