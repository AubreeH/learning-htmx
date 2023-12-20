package main

import (
	"learning-htmx/src/routes"

	"github.com/AubreeH/goApiRouting/routing"

	_ "embed"
)

//go:embed view/pages/index.go.tmpl
var IndexTemplate string

var HomeTemplate string

func main() {
	r := routing.NewRouter(routing.Config{
		Port: 8080,
	})
	r.InitialiseRoutes(routes.SetupRoutes)
	r.Start()
}
