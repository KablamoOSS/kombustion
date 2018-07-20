package cloudformation

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/internal/plugins"
	"github.com/KablamoOSS/kombustion/pkg/parsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

type (
	testYamlTemplateCF struct {
		tstRes  types.ResourceMap
		expcRes types.TemplateObject
		isRes   bool
		log     string
		errMsg  string
	}

	testGenerateYamlStack struct {
		genParams GenerateParams
		errMsg    string
		expect    YamlCloudformation
		file      *os.File
	}
)
