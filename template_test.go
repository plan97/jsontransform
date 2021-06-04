package jsontransform

import (
	"bytes"
	"testing"
	"text/template"
)

func TestParseTemplate(t *testing.T) {
	type args struct {
		name     string
		template string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				name:     "test",
				template: "hello {{.world}}",
			},
			wantErr: false,
		},
		{
			name: "test 2",
			args: args{
				name:     "test",
				template: "hello {{if .world}}",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTemplate(tt.args.name, tt.args.template)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRenderTemplate(t *testing.T) {
	type args struct {
		tmpl *template.Template
		name string
		data string
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				tmpl: template.Must(ParseTemplate("test", "hello {{.name}}")),
				name: "test",
				data: `{"name":"world"}`,
			},
			wantWriter: "hello world",
			wantErr:    false,
		},
		{
			name: "test 2",
			args: args{
				tmpl: template.Must(ParseTemplate("test", "hello {{if .unknown}}?{{end}}")),
				name: "test",
				data: `{"unknown":true}`,
			},
			wantWriter: "hello ?",
			wantErr:    false,
		},
		{
			name: "test 3",
			args: args{
				tmpl: template.Must(ParseTemplate("test", "hello {{if .unknown}}?{{end}}")),
				name: "test",
				data: `unknown data format`,
			},
			wantWriter: "",
			wantErr:    true,
		},
		{
			name: "test 4",
			args: args{
				tmpl: template.Must(ParseTemplate("test", "hello {{len .unknown}}")),
				name: "test",
				data: `{"unknown":null}`,
			},
			wantWriter: "hello ",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := RenderTemplate(tt.args.tmpl, tt.args.name, tt.args.data, writer); (err != nil) != tt.wantErr {
				t.Errorf("RenderTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("RenderTemplate() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
