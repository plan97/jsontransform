package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/plan97/jsontransform"
)

func main() {
	const (
		name                 = "tpl"
		fileMode os.FileMode = 0666
	)

	tplFile := flag.String("tpl", "template.tpl", "File containing the template")
	dataFile := flag.String("data", "data.json", "File containing the template data")
	outFile := flag.String("o", "", "File to write the generated output")
	flag.Parse()

	template, err := os.ReadFile(*tplFile)
	if err != nil {
		log.Fatalln("unable to read template file:", err)
	}

	data, err := os.ReadFile(*dataFile)
	if err != nil {
		log.Fatalln("unable to read template data file:", err)
	}

	buffer := bytes.NewBuffer(nil)
	tmpl, err := jsontransform.ParseTemplate(name, string(template))
	if err != nil {
		log.Fatalln("unable to parse template:", err)
	}

	err = jsontransform.RenderTemplate(tmpl, name, string(data), buffer)
	if err != nil {
		log.Fatalln("unable to render template:", err)
	}

	if *outFile == "" {
		fmt.Print(buffer.String())
	} else {
		err = os.WriteFile(*outFile, buffer.Bytes(), fileMode)
		if err != nil {
			log.Fatalln("unable to write output to file:", err)
		}
	}
}
