package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"

	printer "github.com/KablamoOSS/go-cli-printer"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"github.com/urfave/cli"
)

type YamlConfig struct {
	AWSTemplateFormatVersion string               `yaml:"AWSTemplateFormatVersion,omitempty"`
	Description              string               `yaml:"Description,omitempty"`
	Parameters               ParameterMap         `yaml:"Parameters,omitempty"`
	Mappings                 types.TemplateObject `yaml:"Mappings,omitempty"`
	Conditions               types.TemplateObject `yaml:"Conditions,omitempty"`
	Transform                types.TemplateObject `yaml:"Transform,omitempty"`
	Resources                ResourceMap          `yaml:"Resources"`
	Outputs                  types.TemplateObject `yaml:"Outputs,omitempty"`
}

type CfResource struct {
	Type       string      `yaml:"Type"`
	Properties interface{} `yaml:"Properties"`
	Condition  interface{} `yaml:"Condition,omitempty"`
	Metadata   interface{} `yaml:"Metadata,omitempty"`
	DependsOn  interface{} `yaml:"DependsOn,omitempty"`
}

type CfParameter struct {
	Type                  string      `yaml:"Type"`
	AllowedPattern        interface{} `yaml:"AllowedPattern,omitempty"`
	AllowedValues         interface{} `yaml:"AllowedValues,omitempty"`
	ConstraintDescription interface{} `yaml:"ConstraintDescription,omitempty"`
	Default               interface{} `yaml:"Default,omitempty"`
	Description           interface{} `yaml:"Description,omitempty"`
	MaxLength             interface{} `yaml:"MaxLength,omitempty"`
	MaxValue              interface{} `yaml:"MaxValue,omitempty"`
	MinLength             interface{} `yaml:"MinLength,omitempty"`
	MinValue              interface{} `yaml:"MinValue,omitempty"`
	NoEcho                interface{} `yaml:"NoEcho,omitempty"`
}

type ResourceMap map[string]CfResource
type ParameterMap map[string]CfParameter

var config YamlConfig

///

type intrinsicFn string

func (name intrinsicFn) UnmarshalYAMLTag(t string, out reflect.Value) reflect.Value {
	output := reflect.ValueOf(make(map[interface{}]interface{}))
	output.SetMapIndex(reflect.ValueOf(string(name)), out)
	return output
}

func registerYamlTagUnmarshalers() {
	yaml.RegisterTagUnmarshaler("!Ref", intrinsicFn("Ref"))
	yaml.RegisterTagUnmarshaler("!Base64", intrinsicFn("Fn::Base64"))
	yaml.RegisterTagUnmarshaler("!FindInMap", intrinsicFn("Fn::FindInMap"))
	yaml.RegisterTagUnmarshaler("!Join", intrinsicFn("Fn::Join"))
	yaml.RegisterTagUnmarshaler("!Select", intrinsicFn("Fn::Select"))
	yaml.RegisterTagUnmarshaler("!Split", intrinsicFn("Fn::Split"))
	yaml.RegisterTagUnmarshaler("!Sub", intrinsicFn("Fn::Sub"))
	yaml.RegisterTagUnmarshaler("!And", intrinsicFn("Fn::And"))
	yaml.RegisterTagUnmarshaler("!Equals", intrinsicFn("Fn::Equals"))
	yaml.RegisterTagUnmarshaler("!If", intrinsicFn("Fn::If"))
	yaml.RegisterTagUnmarshaler("!Not", intrinsicFn("Fn::Not"))
	yaml.RegisterTagUnmarshaler("!Or", intrinsicFn("Fn::Or"))
	yaml.RegisterTagUnmarshaler("!GetAtt", intrinsicFn("Fn::GetAtt"))
	yaml.RegisterTagUnmarshaler("!GetAZs", intrinsicFn("Fn::GetAZs"))
	yaml.RegisterTagUnmarshaler("!ImportValue", intrinsicFn("Fn::ImportValue"))
	yaml.RegisterTagUnmarshaler("!Cidr", intrinsicFn("Fn::Cidr"))
}

func init() {
	registerYamlTagUnmarshalers()
}

func getVal(v interface{}, depth int, append string) string {
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Int {
		return fmt.Sprint("\"", v, "\"", append)
	} else if typ == reflect.Bool {
		return fmt.Sprint("\"", v, "\"", append)
	} else if typ == reflect.String {
		return fmt.Sprint("\"", strings.Replace(strings.Replace(v.(string), "\n", "\\n", -1), "\"", "\\\"", -1), "\"", append)
	} else if typ == reflect.Slice {
		return printSlice(v.([]interface{}), depth+1, append)
	} else if typ == reflect.Map {
		return printMap(v.(map[interface{}]interface{}), depth+1, append)
	}

	return "UNKNOWN_TYPE_" + typ.String()
}

func printMap(m map[interface{}]interface{}, depth int, append string) string {
	retStr := ""
	prepend := ""
	hasSubmaps := true

	if depth == 1 {
		for k, v := range m {
			retStr += "\n" + k.(string) + ": " + getVal(v, depth, "") + ","
		}
	} else {
		hasSubmaps = false
		for k, _ := range m {
			typ := reflect.TypeOf(k).Kind()
			if typ != reflect.Int && typ != reflect.String {
				hasSubmaps = true
			}
		}
		if hasSubmaps {
			retStr += "map[interface{}]interface{}{"
		} else {
			retStr += "map[string]interface{}{"
		}

		for k, v := range m {
			prepend = ""
			val := getVal(v, depth, "")
			typ := reflect.TypeOf(v).Kind()
			key := getVal(k.(string), depth+1, "")

			if typ == reflect.Int || typ == reflect.String {
				// Check to see if it's a ref/getatt and prepend custom resource name
				if k.(string) == "Ref" || k.(string) == "Fn::GetAtt" {
					for resName := range config.Resources {
						if val == ("\"" + resName + "\"") {
							prepend = "name + "
						}
					}
					for paramName := range config.Parameters {
						if val == ("\"" + paramName + "\"") {
							return "param" + paramName + append
						}
					}
				}
			}
			retStr += "\n" + key + ": " + prepend + val + ","
		}
		retStr = retStr[:len(retStr)-1] // trim last comma
		retStr += "}" + append
	}

	return retStr
}

func printSlice(slc []interface{}, depth int, append string) string {
	retStr := "[]interface{}{\n"

	for _, v := range slc {
		retStr += getVal(v, depth, ",\n")
	}

	retStr += "}"

	return retStr
}

func main() {
	app := cli.NewApp()
	app.Name = "genplugin"
	app.Usage = "genplugin --inputfile example.yaml --pluginname myplugin --resourcename ExampleType --resourcetype Kombustion::Example::Type"
	app.Action = generatePlugin
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "inputfile",
			Usage: "native CloudFormation file (only YAML is currently supported)",
		},
		cli.StringFlag{
			Name:  "pluginname",
			Usage: "the name of the plugin package",
		},
		cli.StringFlag{
			Name:  "resourcename",
			Usage: "the name of the resource that the stack represents",
		},
		cli.StringFlag{
			Name:  "resourcetype",
			Usage: "the value that should be used for the Type field (typically in the form Aaa::Bbb::Ccc)",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		// TODO: Make this error more helpful
		printer.Fatal(
			err,
			"",
			"",
		)
	}

}

func generatePlugin(c *cli.Context) error {
	inputfile := c.String("inputfile")
	pluginname := c.String("pluginname")
	resname := c.String("resourcename")
	restype := c.String("resourcetype")

	config = GetCloudformationClientn(inputfile)
	resources := generateResources(config, resname)
	mainfile := generateMain(pluginname, restype, resname, config)

	os.MkdirAll("../plugins/"+pluginname+"/resources/", 0744)
	writeOutput("../plugins/"+pluginname+"/resources/"+strings.ToLower(resname)+".go", resources)
	writeOutput("../plugins/"+pluginname+"/plugin.go", mainfile)

	cmd := exec.Command("go", "build", "-buildmode", "plugin")
	cmd.Dir = fmt.Sprintf("../plugins/%v", pluginname)
	cmdOut, err := cmd.CombinedOutput()

	fmt.Println(string(cmdOut))

	return err
}

func generateResources(config YamlConfig, resname string) []byte {
	buf := bytes.NewBufferString("")
	needsResourcesImport := false

	for _, res := range config.Resources {
		if res.Metadata != nil {
			needsResourcesImport = true
		}
	}
	if len(config.Parameters) > 0 {
		needsResourcesImport = true
	}

	writeLine(buf, "package resources\n"+
		"\n"+
		"import (\n"+
		"	\"github.com/KablamoOSS/kombustion/internal/plugins\"\n")
	if needsResourcesImport {
		writeLine(buf, "   \"github.com/KablamoOSS/kombustion/pluginParsers/resources\"\n"+
			"   yaml \"gopkg.in/yaml.v2\"\n")
	}
	writeLine(buf, ")\n"+
		"\n")

	if len(config.Parameters) > 0 {
		writeLine(buf, "type "+resname+"Config struct {\n"+
			"	Properties struct {\n")
		for paramName, param := range config.Parameters {
			if strings.HasPrefix(param.Type, "List<") || param.Type == "CommaDelimitedList" {
				writeLine(buf, "		"+paramName+" []interface{} `yaml:\""+paramName+",omitempty\"`\n")
			} else {
				writeLine(buf, "		"+paramName+" *string `yaml:\""+paramName+",omitempty\"`\n")
			}
		}
		writeLine(buf, "	} `yaml:\"Properties\"`\n"+
			"}\n\n")
	}

	writeLine(buf, "func Parse"+resname+"(name string, data string) (cf types.TemplateObject, err error) {\n"+
		"	// create a group of objects (each to be validated)\n"+
		"	cf = make(types.TemplateObject)\n"+
		"\n")

	if len(config.Parameters) > 0 {
		writeLine(buf, "// Parse the config data\n"+
			"var config "+resname+"Config\n"+
			"if err = yaml.Unmarshal([]byte(data), &config); err != nil {\n"+
			"  return\n"+
			"}\n"+
			"\n"+
			"// validate the config\n"+
			"config.Validate()\n\n// defaults\n")
		for paramName, param := range config.Parameters {
			if strings.HasPrefix(param.Type, "List<") || param.Type == "CommaDelimitedList" {
				writeLine(buf, "param"+paramName+" := []interface{}{}\n"+
					"if len(config.Properties."+paramName+") > 0 {\n"+
					"  param"+paramName+" = config.Properties."+paramName+"\n"+
					"}\n\n")
			} else {
				defaultVal := ""
				if param.Default != nil {
					defaultVal = fmt.Sprintf("%v", param.Default) // ensure it's a string if int is given
				}
				writeLine(buf, "param"+paramName+" := \""+defaultVal+"\"\n"+
					"if config.Properties."+paramName+" != nil {\n"+
					"  param"+paramName+" = *config.Properties."+paramName+"\n"+
					"}\n\n")
			}
		}
	}

	for resName, res := range config.Resources {
		typeParts := strings.Split(res.Type, "::")

		if res.Metadata != nil || res.Condition != nil || res.DependsOn != nil {
			writeLine(buf, "resource", resName, " := resources.New", typeParts[1], typeParts[2], "(\n")
		} else {
			writeLine(buf, "cf[name+\"", resName, "\"] = resources.New", typeParts[1], typeParts[2], "(\n")
		}
		writeLine(buf, "resources.", typeParts[1], typeParts[2], "Properties{")

		writeLine(buf, getVal(res.Properties, 0, ""))

		writeLine(buf, "\n},\n")
		writeLine(buf, ")\n\n")

		if res.Metadata != nil {
			writeLine(buf, "resource", resName, ".Metadata = ", getVal(res.Metadata, 1, ""), "\n\n")
		}

		if res.Condition != nil {
			writeLine(buf, "resource", resName, ".Condition = ", getVal(res.Condition, 1, ""), "\n\n")
		}

		if res.DependsOn != nil {
			writeLine(buf, "resource", resName, ".DependsOn = ", getVal(res.DependsOn, 1, ""), "\n\n")
		}

		if res.Metadata != nil || res.Condition != nil || res.DependsOn != nil {
			writeLine(buf, "cf[name+\"", resName, "\"] = resource", resName, "\n\n")
		}
	}

	writeLine(buf, "return\n}\n")

	if len(config.Parameters) > 0 {
		writeLine(buf, "// Validate - input Config validation\n"+
			"func (this "+resname+"Config) Validate() {\n"+
			"    \n"+
			"}\n")
	}

	formatted, err := format.Source(buf.Bytes()) // lint

	if err != nil {
		fmt.Println("ERROR: Could not format!")
		return buf.Bytes()
	}

	return formatted
}

func generateMain(pluginname string, restype string, resname string, config YamlConfig) []byte {
	buf := bytes.NewBufferString("")

	writeLine(buf, "package main\n"+
		"\n"+
		"import (\n"+
		"	\"github.com/KablamoOSS/kombustion/internal/plugins/"+pluginname+"/resources\"\n"+
		"	\"github.com/KablamoOSS/kombustion/internal/plugins\"\n"+
		")\n"+
		"\n"+
		"var Resources = map[string]plugins.ParserFunc{\n"+
		"	\""+restype+"\": resources.Parse"+resname+",\n"+
		"}\n"+
		"\n"+
		"var Outputs = map[string]plugins.ParserFunc{}\n"+
		"\n"+
		"var Mappings = map[string]plugins.ParserFunc{}\n"+
		"\n"+
		"func main() {}\n")

	return buf.Bytes()
}

func GetCloudformationClientn(cfPath string) YamlConfig {
	// load the config file
	var data []byte
	data, err := ioutil.ReadFile(cfPath)
	checkError(err)

	var config YamlConfig
	err = yaml.Unmarshal(data, &config)
	checkError(err)

	return config
}

func writeOutput(path string, output []byte) {
	err := ioutil.WriteFile(path, output, 0644)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		// TODO: Make this error more helpful
		printer.Fatal(
			err,
			"",
			"",
		)
	}
}

func writeLine(buf *bytes.Buffer, ss ...string) {
	for _, s := range ss {
		buf.WriteString(s)
	}
}
