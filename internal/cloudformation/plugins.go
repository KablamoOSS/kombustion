package cloudformation

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"plugin"
	"runtime"

	"github.com/KablamoOSS/kombustion/types"
	log "github.com/sirupsen/logrus"
)

const RepositoryPath = "https://downloads.kombustion.io/plugins"

func getPluginDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	plugindir := fmt.Sprintf("%s/.kombustion/plugins", usr.HomeDir)
	os.MkdirAll(plugindir, 0744)

	return plugindir
}

func deletePlugin(pluginname string) error {
	filename := fmt.Sprintf("%s/%s.so", getPluginDir(), pluginname)
	err := os.Remove(filename)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func downloadPlugin(pluginname string) error {
	url := ""
	if runtime.GOOS == "darwin" {
		url = fmt.Sprintf("%s/MacOS/64/%s.so", RepositoryPath, pluginname)
	} else if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
		url = fmt.Sprintf("%s/Linux/64/%s.so", RepositoryPath, pluginname)
	} else {
		log.Fatal("Unsupported operating system or architecture: ", runtime.GOOS, runtime.GOARCH)
	}

	filename := fmt.Sprintf("%s/%s.so", getPluginDir(), pluginname)

	output, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func loadPlugins() (resources, outputs, mappings map[string]types.ParserFunc) {
	resources, outputs, mappings = make(map[string]types.ParserFunc), make(map[string]types.ParserFunc), make(map[string]types.ParserFunc)

	pluginBaseDir := ""
	if len(os.Getenv("PLUGINS")) > 0 {
		pluginBaseDir = os.Getenv("PLUGINS")
	} else {
		pluginBaseDir = getPluginDir()
	}

	pluginDirectories, err := ioutil.ReadDir(pluginBaseDir)
	if err != nil {
		// log.Println("WARNING: ", err)
		return
	}

	for _, d := range pluginDirectories {
		if !d.IsDir() {
			// load all plugins in the base dir
			loadPlugin(d.Name(), pluginBaseDir, resources, outputs, mappings)
			continue
		}

		pluginDir := filepath.Join(pluginBaseDir, d.Name())
		plugins, err := ioutil.ReadDir(pluginDir)
		if err != nil {
			continue
		}

		// load all plugins in each sub dir
		for _, f := range plugins {
			loadPlugin(f.Name(), pluginDir, resources, outputs, mappings)
		}
	}
	return
}

func loadPlugin(filename, pluginDir string, resources, outputs, mappings map[string]types.ParserFunc) {
	if filepath.Ext(filename) != ".so" && filepath.Ext(filename) != ".dll" && filepath.Ext(filename) != ".dylib" {
		return
	}

	pluginPath := filepath.Join(pluginDir, filename)

	log.Info("Using plugin: ", pluginPath)
	p, err := plugin.Open(pluginPath)

	if err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"err":      err,
		}).Warn("error reading plugin file")
		return
	}

	r, err := p.Lookup("Resources")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"err":      err,
		}).Warn("error reading resource plugin")
	}
	o, err := p.Lookup("Outputs")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"err":      err,
		}).Warn("error reading resource plugin")
	}
	m, err := p.Lookup("Mappings")
	if err != nil {
		log.WithFields(log.Fields{
			"filename": filename,
			"err":      err,
		}).Warn("error reading resource plugin")
	}

	for k, v := range *r.(*map[string]types.ParserFunc) {
		if _, ok := resources[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"resource": k,
			}).Warn("duplicate resource definition for resource")
		} else {
			resources[k] = v
		}
	}
	for k, v := range *o.(*map[string]types.ParserFunc) {
		if _, ok := outputs[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"output": k,
			}).Warn("duplicate output definition for output")
		} else {
			outputs[k] = v
		}
	}
	for k, v := range *m.(*map[string]types.ParserFunc) {
		if _, ok := mappings[k]; ok { // Check for duplicates
			log.WithFields(log.Fields{
				"mapping": k,
			}).Warn("duplicate mapping definition for mapping")
		} else {
			mappings[k] = v
		}
	}
	return
}
