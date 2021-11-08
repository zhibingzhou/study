package router

import "poke/api"

func HelloWorldInit() {

	var hello_router = api.ApiGroupApp.ExampleApi
	Router.Get("/hello", hello_router.HelloWorld.Show)

}
