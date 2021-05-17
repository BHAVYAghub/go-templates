package main

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"log"
)

type Data struct{
	Name string
	Desc string
	Fields []Field
}
type Field struct{
	Name string
	TypeName string
}
func main(){
	data:=[]Data{
		{
			Name: "Struct_1",
			Desc: "Details of struct 1",
			Fields: []Field{
				{"Id","int"},
				{"field_1","string"},
				{"Active","boolean"},
			},
		},
		{
			Name: "Struct_2",
			Desc: "Details of struct 2",
			Fields: []Field{
				{"Id","int"},
				{"field_2","string"},
				{"Active","boolean"},
			},
		},
	}
	tmpl,err:=template.ParseFiles("struct.tmpl")
	if err!=nil{
		log.Println("could not parse the file")
	}
	var processed  bytes.Buffer
	err=tmpl.Execute(&processed,data)
	if err!=nil{
		log.Println(err)
		panic("unable to execute")
	}
	formatted,err:=format.Source(processed.Bytes())
	if err!=nil{
		panic("unable to format")
	}
	fmt.Println(string(formatted))

}
