package controllers

import (
	"fmt"
	"net/url"

	"github.com/AubreeH/goApiRouting/routing"
)

var pathPageNames = map[string]string{
	"/":         "Index",
	"/home":     "Home",
	"/messages": "Messages",
}

func getPageName(path string) string {
	if name, ok := pathPageNames[path]; ok {
		return name
	}

	return ""
}

func getCurrentPath(c *routing.Context) (string, error) {
	currentUrl := c.Request.Header.Get("Hx-Current-Url")
	u, err := url.Parse(currentUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse current url: %w", err)
	}

	return u.Path, nil
}
