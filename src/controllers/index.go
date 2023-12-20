package controllers

import (
	"learning-htmx/view"
	"path"
	"strings"

	"github.com/AubreeH/goApiRouting/routing"
)

func Index(r *routing.Context) routing.Response {
	requestPath := r.Request.URL.Path

	if strings.HasPrefix(requestPath, "/page") {
		return routing.Response{
			Status: 301,
			Headers: map[string]string{
				"Location": "/error/not-found.html",
			},
		}
	}

	return routing.Response{
		Headers: map[string]string{
			"HX-Trigger": "page-changed-event",
		},
		Template: view.IndexTemplate,
		Body: map[string]interface{}{
			"Title":     getPageName(requestPath),
			"TitleMeta": "Learning HTMX",
			"Stylesheets": []string{
				"main.css",
				"index__styles.css",
			},
			"RequestPath": requestPath,
			"PagePath":    path.Join("page", requestPath),
		},
	}
}
