package view

import (
	_ "embed"
	"html/template"
)

//go:embed pages/index.go.tmpl
var indexTemplateStr string
var IndexTemplate = template.Must(template.New("index").Parse(indexTemplateStr))

//go:embed pages/home.go.tmpl
var homeTemplateStr string
var HomeTemplate = template.Must(template.New("index").Parse(homeTemplateStr))

//go:embed pages/messages.go.tmpl
var messagesTemplateStr string
var MessagesTemplate = template.Must(template.New("index").Parse(messagesTemplateStr))
