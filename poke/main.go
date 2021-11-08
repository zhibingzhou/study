package main

import (
	"fmt"
	_ "poke/initialize"
	"poke/router"
)

func main() {

	fmt.Println("this is fiber!!")

	router.Router.Listen(":80")
}
