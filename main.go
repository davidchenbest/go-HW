package main

import (
	"html/template"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/content.html"))
	tmpl.Execute(os.Stdout, "World")

}
