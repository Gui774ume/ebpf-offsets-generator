package utils

import "html/template"

// Unescape - Unescape value in a template
func Unescape(s string) template.HTML {
	return template.HTML(s)
}
