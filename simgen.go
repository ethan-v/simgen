package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

/**
 * This is a simple command-line tool that generates a file from a template and JSON data.
 * It is used to generate the form.html file from the form.html.template and form.json files.
 * Usage: simgen -template=form.html.template -data=form.json -output=form.html
 * Example: go run simgen.go -template=./templates/form/form.template.html -data=./templates/form/data.json -output=form.html
 */
func main() {
	// define command-line flags
	templateFile := flag.String("template", "", "path to template file")
	dataFile := flag.String("data", "", "path to data file")
	outputFile := flag.String("output", "", "path to output file")
	flag.Parse()

	// validate command-line flags
	if *templateFile == "" {
		fmt.Println("Error: template file not specified")
		os.Exit(1)
	}
	if *dataFile == "" {
		fmt.Println("Error: data file not specified")
		os.Exit(1)
	}
	if *outputFile == "" {
		fmt.Println("Error: output file not specified")
		os.Exit(1)
	}

	// read template file
	templateContent, err := ioutil.ReadFile(*templateFile)
	if err != nil {
		fmt.Println("Error reading template file:", err)
		os.Exit(1)
	}

	// create template object
	tmpl, err := template.New("template").Parse(string(templateContent))
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}

	// read data file
	dataContent, err := ioutil.ReadFile(*dataFile)
	if err != nil {
		fmt.Println("Error reading data file:", err)
		os.Exit(1)
	}

	// parse JSON data
	var data map[string]interface{}
	err = json.Unmarshal(dataContent, &data)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		os.Exit(1)
	}

	// execute template with data
	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer output.Close()

	err = tmpl.Execute(output, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		os.Exit(1)
	}

	fmt.Println("Output written to", *outputFile)
}
