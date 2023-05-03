package main

import (
	"miniproject/config"
	"miniproject/routes"
)

func main() {
	config.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
