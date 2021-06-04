[{{range $i, $v := .}}{{if $i}},{{end}}
    {
        "hello": "{{.name}}",
        "data": {{if .data}}[{{range $ii, $vv := .data}}{{if $ii}}, {{end}}{"value": "{{$vv}}"}{{end}}]{{else}}null{{end}}
    }{{end}}
]