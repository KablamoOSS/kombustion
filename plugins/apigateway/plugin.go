// +build plugin

package main

import (
	"github.com/KablamoOSS/kombustion/plugins/apigateway/resources"
	"github.com/KablamoOSS/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
	"Kombustion::ApiGateway::API": resources.ParseApiGateway,
}

var Outputs = map[string]types.ParserFunc{}

var Mappings = map[string]types.ParserFunc{}

func main() {}
