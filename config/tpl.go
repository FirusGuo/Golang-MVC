package config

import (
	"html/template"
)

// TPL init for templates
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
