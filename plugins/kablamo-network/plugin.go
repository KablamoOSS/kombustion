// +build plugin

package main

import (
	"github.com/KablamoOSS/kombustion/types"
	//"github.com/KablamoOSS/kombustion/plugins/kablamo-network/mappings"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-network/outputs"
	"github.com/KablamoOSS/kombustion/plugins/kablamo-network/resources"
)

var Resources = map[string]types.ParserFunc{
	"Kablamo::Network::VPC": resources.ParseNetworkVPC,
}

var Outputs = map[string]types.ParserFunc{
	"Kablamo::Network::VPC": outputs.ParseNetworkVPC,
}

var Mappings = map[string]types.ParserFunc{}

func main() {}
