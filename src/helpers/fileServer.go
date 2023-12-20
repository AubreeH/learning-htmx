package helpers

import (
	"path"
	"strings"

	"github.com/AubreeH/goApiRouting/routing"
)

func FileServer(routePrefix, dirPath string) (string, func(c *routing.Context) routing.Response) {
	if len(routePrefix) == 0 || routePrefix[0] != '/' {
		routePrefix = "/" + routePrefix
	}

	return routePrefix + "/*", func(c *routing.Context) routing.Response {
		requestPath := strings.TrimPrefix(c.Request.URL.Path, routePrefix)
		filePath := path.Join(dirPath, requestPath)

		if ok, err := FileExists(filePath); err != nil {
			return routing.Response{
				Status: 500,
				Body:   "Internal Server Error",
			}
		} else if !ok {
			return routing.Response{
				Status: 404,
				Body:   "Not Found",
			}
		}

		return routing.Response{
			Body: filePath,
			Type: routing.FileResponse,
		}
	}
}
