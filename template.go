/*
Package jsontransform provides parsers and renderers using JSON data.

Parsed templates can make use of functions from https://github.com/Masterminds/sprig.
*/
package jsontransform

import (
	"fmt"
	"io"
	t "text/template"

	sprig "github.com/Masterminds/sprig/v3"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseTemplate creates a new template parser with the input name and template
// and adds functions from https://github.com/Masterminds/sprig.
func ParseTemplate(name, template string) (*t.Template, error) {
	tmpl, err := t.New(name).Funcs(sprig.TxtFuncMap()).Parse(template)
	if err != nil {
		return nil, fmt.Errorf("unable to parse template: %w", err)
	}
	return tmpl, nil
}

// RenderTemplate uses the template parser with the input name
// to render the output using the provided data (JSON).
// The output will be written to the writer.
func RenderTemplate(tmpl *t.Template, name, data string, writer io.Writer) error {
	var d interface{}
	err := json.Unmarshal([]byte(data), &d)
	if err != nil {
		return fmt.Errorf("unable to unmarshal json: %w", err)
	}

	err = tmpl.ExecuteTemplate(writer, name, d)
	if err != nil {
		return fmt.Errorf("error while executing template: %w", err)
	}
	return nil
}
