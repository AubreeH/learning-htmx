package controllers

import (
	"path"
	"strings"
	"text/template"

	"github.com/AubreeH/goApiRouting/routing"
)

var tmplSidePanel = template.Must(template.ParseFiles("view/component/sidePanel.go.tmpl"))

func SidePanel(c *routing.Context) routing.Response {
	path, err := getCurrentPath(c)
	if err != nil {
		return routing.Response{
			Status: 500,
			Body:   "Internal Server Error",
		}
	}

	return routing.Response{
		Template: tmplSidePanel,
		Body: []interface{}{
			sidePanelItem(path, "Home", "/home"),
			sidePanelItem(path, "Messages", "/messages"),
		},
	}
}

func sidePanelItem(currentPath string, title, itemPath string) map[string]interface{} {
	return map[string]interface{}{
		"Title":         title,
		"Path":          itemPath,
		"ComponentPath": path.Join("page", itemPath),
		"Active":        isSidePanelItemActive(currentPath, itemPath),
	}
}

func isSidePanelItemActive(currentPath string, path string) bool {
	return strings.HasPrefix(currentPath, path)
}
