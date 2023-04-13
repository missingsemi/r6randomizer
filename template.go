package main

import (
	_ "embed"
	"html/template"
	"io"
)

//go:embed template.html.tmpl
var rawTemplate string
var tmpl = template.Must(template.New("").Parse(rawTemplate))

func RenderTemplate(w io.Writer, l Loadout) error {
	return tmpl.Execute(w, l)
}
