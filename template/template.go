package template

import (
	"bytes"
	"text/template"
)

type Template struct {
	tmpl *template.Template
}

func (t *Template) RenderMessage(data interface{}) (string, error) {
	var tpl bytes.Buffer
	if err := t.tmpl.ExecuteTemplate(&tpl, "message", data); err != nil {
		return "", err
	}
	return tpl.String(), nil
}

func NewTemplate(path string) (*Template, error) {
	template, err := template.New("").ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return &Template{
		tmpl: template,
	}, nil
}
