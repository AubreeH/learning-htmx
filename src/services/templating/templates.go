package templating

import (
	"html/template"
)

type Template struct {
	FilePath string
	Name     string
	tmpl     *template.Template
}

type Params struct {
	Styles   []string
	Scripts  []string
	Children []string
}

var templates = map[string]Template{}
