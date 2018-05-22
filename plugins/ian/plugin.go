// +build plugin

package main

import (
	"github.com/KablamoOSS/kombustion/plugins/ian/mappings"
	"github.com/KablamoOSS/kombustion/plugins/ian/outputs"
	"github.com/KablamoOSS/kombustion/plugins/ian/resources"
	"github.com/KablamoOSS/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
	"Ian::Databases::Aurora":          resources.ParseAurora,
	"Ian::WebApplications::Wordpress": resources.ParseWordpress,
	"Ian::Meta::Attributes":           resources.ParseMetaAttributes,
}

var Outputs = map[string]types.ParserFunc{
	"Ian::Databases::Aurora":          outputs.ParseAurora,
	"Ian::WebApplications::Wordpress": outputs.ParseWordpress,
}

var Mappings = map[string]types.ParserFunc{
	"Ian::WebApplications::Wordpress": mappings.ParseWordpress,
	"Ian::Meta::Attributes":           mappings.ParseMetaAttributes,
}

var Help = types.PluginHelp{
	Description: "A sample plugin",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Ian::Databases::Aurora",
			Description: "Creates an Aurora cluster with the desired number of instances.",
			Config:      resources.AuroraConfig{},
		},
	},
}

func main() {}
