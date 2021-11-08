package api

import "poke/api/example"

type ApiGroup struct {
	ExampleApi example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
