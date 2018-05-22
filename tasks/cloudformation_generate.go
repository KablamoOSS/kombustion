package tasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KablamoOSS/kombustion/cloudformation"
	yaml "github.com/KablamoOSS/yaml"
	"github.com/urfave/cli"
)

var Generate_Flags = []cli.Flag{
	cli.StringFlag{
		Name:  "format, f",
		Usage: "cf output format (only yaml is supported)",
		Value: "yaml",
	},
	cli.StringFlag{
		Name:  "env",
		Usage: "environment config to use from ./config/environment.yaml",
	},
	cli.StringFlag{
		Name:  "envFile",
		Usage: "path to the environment.yaml file",
	},
	cli.StringSliceFlag{
		Name:  "param, p",
		Usage: "cloudformation parameters. eg. ( -p Env=dev -p BucketName=test )",
	},
	cli.BoolFlag{
		Name:  "noBaseOutputs, b",
		Usage: "disable generation of outputs for Base AWS types",
	},
}

func Generate(c *cli.Context) {
	output, _ := generateTemplate(c)
	writeOutput(c, output)
}

func generateTemplate(c *cli.Context) ([]byte, cloudformation.YamlCloudformation) {
	switch c.String("format") {
	case "yaml":
		output, cf := generateYamlTemplate(c)
		return output, cf
	default:
		log.Fatal("Format not supported: ", c.String("format"))
	}
	return []byte{}, cloudformation.YamlCloudformation{}
}

func generateYamlTemplate(c *cli.Context) ([]byte, cloudformation.YamlCloudformation) {
	paramMap := getParamMap(c)

	cf, err := cloudformation.GenerateYamlStack(
		cloudformation.GenerateParams{
			Filename:           c.Args().Get(0),
			EnvFile:            c.String("envFile"),
			Env:                c.String("env"),
			DisableBaseOutputs: c.Bool("noBaseOutputs"),
			ParamMap:           paramMap,
		})
	checkError(err)
	output, err := yaml.Marshal(cf)
	checkError(err)
	return output, cf
}

func writeOutput(c *cli.Context, output []byte) {
	filename := filepath.Base(c.Args().Get(0))
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	path := fmt.Sprint("compiled/", basename, ".yaml")
	os.Mkdir("./compiled", 0744)
	err := ioutil.WriteFile(path, output, 0644)
	checkError(err)
}
