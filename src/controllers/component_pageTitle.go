package controllers

import (
	"github.com/AubreeH/goApiRouting/routing"
)

func PageTitle(c *routing.Context) routing.Response {
	path, err := getCurrentPath(c)
	if err != nil {
		return routing.Response{
			Status: 500,
			Body:   "Internal Server Error",
		}
	}

	return routing.Response{
		Status: 200,
		Body:   getPageName(path),
	}
}
