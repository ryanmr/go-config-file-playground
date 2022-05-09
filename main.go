package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version int64   `yaml:"version"`
	Routes  []Route `yaml:"routes"`
}

type Route struct {
	Route string `yaml:"route"`
	App   string `yaml:"app"`
	Team  string `yaml:"team"`
	Desc  string `yaml:"desc"`
	Repo  string `yaml:"repo"`
}

func main() {

	data, err := os.ReadFile("config.yml")
	check(err)
	fmt.Println("file output")
	fmt.Print(string(data))

	if err != nil {
		log.Printf("read file failed #%v ", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	fmt.Println("struct output")
	fmt.Printf("%+v", config)
	fmt.Println()

	const placeholder = `
### config file version: {{.Version}}

{{range $val := .Routes}}
### Route: {{$val.Route}}
### Desc: {{$val.Desc}}
### Team: {{$val.Team}}
### App: {{$val.App}}
{{if $val.Repo}}### Repo: {{$val.Repo}}{{end}}
{{end}}

### end of file
	`

	tmpl, err := template.New("config").Parse(placeholder)

	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, config); err != nil {
		log.Fatalf("unmarshal: %v", err)
	}

	fmt.Println(tpl.String())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
