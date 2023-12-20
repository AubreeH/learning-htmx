package controllers

import (
	"learning-htmx/view"

	"github.com/AubreeH/goApiRouting/routing"
)

func Messages(c *routing.Context) routing.Response {
	return routing.Response{
		Body:     []string{"abc", "def", "ghi"},
		Template: view.MessagesTemplate,
	}
}
