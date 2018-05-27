package main

import (
	"github.com/KablamoOSS/kombustion/types"
	//"bitbucket.org/kablamo-dev/kombustion/plugins/kablamo-securitygroups/mappings"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-server/outputs"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-server/resources"
)

var Resources = map[string]types.ParserFunc{
	"Kablamo::EC2::Server": resources.ParseEC2Server,
}

var Outputs = map[string]types.ParserFunc{
	"Kablamo::EC2::Server": outputs.ParseEC2Server,
}

var Mappings = map[string]types.ParserFunc{}

var Help = types.PluginHelp{
	Description: "Helper function for Kablamo VPC",
	TypeMappings: []types.TypeMapping{
		{
			Name:        "Kablamo::Network::VPC",
			Description: "Creates a complete VPC network with subnets, route tables, routes & NACL's",
			Config:      resources.EC2ServerConfig{},
		},
	},
}

func main() {}
