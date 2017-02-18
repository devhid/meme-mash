package main

import (
	"log"
	"os"
	"text/template"
)

func view() {
	tpl, err := template.ParseFiles("view.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, getImageArray())
}
