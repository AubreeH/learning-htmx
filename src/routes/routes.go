package routes

import (
	"learning-htmx/src/controllers"
	"learning-htmx/src/helpers"

	"github.com/AubreeH/goApiRouting/routing"
)

func setHxTriggerMiddleware(header string) func(c *routing.Context, respond func(routing.Response)) bool {
	return func(c *routing.Context, respond func(routing.Response)) bool {
		c.Writer.Header().Set("HX-Trigger", "page-changed-event")
		return true
	}
}

func SetupRoutes(api routing.BaseApi) {
	api.Get("*", controllers.Index)

	api.Any(helpers.FileServer("/assets/", "view/assets"))
	api.Any(helpers.FileServer("/common/", "view/common"))
	api.Any(helpers.FileServer("/error/", "view/error"))
	api.Any(helpers.FileServer("/", "view/root"))

	api.Group("/page", routing.WithMiddleware(setHxTriggerMiddleware("page-changed-event")), func(api routing.BaseApi) {
		api.Get("/messages", controllers.Messages)
		api.Get("/home", controllers.Home)
	})

	api.Group("/component", api.NoMiddleware(), func(api routing.BaseApi) {
		api.Get("/side-panel", controllers.SidePanel)
		api.Get("/page-title", controllers.PageTitle)
	})
}
