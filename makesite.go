package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

type content struct {
	Content string
}

func main() {
	firstPost, _ := ioutil.ReadFile("first-post.text")
	t := template.Must(template.New("template.tmpl").ParseFiles("first-post.html"))
	err := t.Execute(os.Stdout, firstPost)
	if err != nil {
		panic(err)
	}
}
