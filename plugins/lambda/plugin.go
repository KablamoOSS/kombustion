// +build plugin

package main

import (
	"github.com/KablamoOSS/kombustion/plugins/lambda/resources"
	"github.com/KablamoOSS/kombustion/types"
)

var Resources = map[string]types.ParserFunc{
	"Kombustion::Lambda::Permission": resources.ParseLambdaPermission,
}

var Outputs = map[string]types.ParserFunc{}

var Mappings = map[string]types.ParserFunc{}

func main() {}
