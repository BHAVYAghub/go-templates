package main

import (
	"html/template"
	"os"
	"strings"
)

func main(){
	tmplString:="Hello, {{.}}"
	name:="Donald"
	tmpl:=template.Must(template.New("hello").Parse(tmplString))
	err:=tmpl.Execute(os.Stdout,name)
	if err!=nil{
		panic("could not execute tmpl hello")
	}

	tmplString="{{ if . }}Hello{{ end }}"
	tmpl=template.Must(template.New("hello2").Parse(tmplString))
	err=tmpl.Execute(os.Stdout,1)
	if err!=nil{
		panic("could not execute tmpl hello2")
	}

	tmplString="{{ range .}}Hello, {{.}}\n{{end}}"
	tmpl=template.Must(template.New("range").Parse(tmplString))
	names:=[]string{"Donald","Bob","Brodie"}
	err=tmpl.Execute(os.Stdout,names)
	if err!=nil{
		panic("could not execute tmpl hello2")
	}

	tmplString="{{ range .}}Hello, {{ toUpper .}}\n{{end}}"
	funcMaps:=map[string]interface{}{
		"toUpper": strings.ToUpper,
	}
	tmpl=template.Must(template.New("funcs").Funcs(funcMaps).Parse(tmplString))
	err=tmpl.Execute(os.Stdout,names)
}