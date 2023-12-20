package controllers

import (
	"learning-htmx/view"

	"github.com/AubreeH/goApiRouting/routing"
)

func Home(c *routing.Context) routing.Response {
	return routing.Response{
		Template: view.HomeTemplate,
	}
}
