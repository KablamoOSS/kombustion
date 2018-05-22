// +build plugin

package main

import (
	"github.com/KablamoOSS/kombustion/types"
	//"github.com/KablamoOSS/kombustion/plugins/kablamo-securitygroups/mappings"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-securitygroups/outputs"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-securitygroups/resources"
)

var Resources = map[string]types.ParserFunc{
	"Kablamo::Network::SecurityGroups": resources.ParseNetworkSecurityGroups,
}

var Outputs = map[string]types.ParserFunc{
	"Kablamo::Network::SecurityGroups": outputs.ParseNetworkSecurityGroups,
}

var Mappings = map[string]types.ParserFunc{}

func main() {}
