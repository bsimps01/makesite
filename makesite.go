package main

import (
	//"flag"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// type entry struct {
// 	Name string
// 	Done bool
// }

// type ToDo struct {
// 	User string
// 	List []entry
// }

type content struct {
	Content string
}

func readFile(templateName string) string {
	fileContents, err := ioutil.ReadFile(templateName)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func buildTemplate(filename string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(filename))

	var err error
	err = t.Execute(os.Stdout, content{Content: data})
	if err != nil {
		panic(err)
	}
}
func addExtHTML(filename string) string {
	ext := ".html"
	return strings.Split(filename, ".")[0] + ext
}

func writeTemplateToFile(templateName string, data string) {
	t := template.Must(template.New("template.tmpl").ParseFiles(templateName))

	file := addExtHTML(data)
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, content{Content: readFile(data)})
	if err != nil {
		panic(err)
	}
}

func isTxtFile(filename string) bool {
	if strings.Contains(filename, ".") {
		return strings.Split(filename, ".")[1] == "txt"
	}
	return false
}

func saveToFile() {
	firstPost, _ := ioutil.ReadFile("first-post.text")
	t := template.Must(template.New("template.tmpl").ParseFiles("first-post.html"))
	err := t.Execute(os.Stdout, firstPost)
	if err != nil {
		panic(err)
	}
	fileParse := flag.String("file", "", "text file will change to an html file")
	directory := flag.String("dir", "", "search files in the directory")
	if *directory != "" {
		textFiles, err := ioutil.ReadDir(*directory)
		if err != nil {
			panic(err)
		}
		var numFiles int
		for _, file := range textFiles {
			filename := file.Name()
			if isTxtFile(filename) == true {
				buildTemplate("template.tmpl", readFile(filename))
				writeTemplateToFile("template.tmpl", filename)
				numFiles++

			}
		}
	}

	if *fileParse != "" {
		buildTemplate("template.tmpl", readFile(*fileParse))
		writeTemplateToFile("template.tmpl", *fileParse)

	} else {
		buildTemplate("template.tmpl", readFile("first-post.txt"))
		writeTemplateToFile("template.tmpl", "test.txt")
	}
}

func main() {
	saveToFile()
}
