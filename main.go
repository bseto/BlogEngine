package main

import (
	"github.com/bseto/BlogEngine/logger"
	"html/template"
	"net/http"
	"path/filepath"
)

type GeneralPage struct {
	ActiveTab string
}

type Article struct {
	Title      string
	Body       string
	CreateDate string
}

func ListArticles() {

}

func Home(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Home"}
	logger.Log("Inside Home")
	RenderTemplate(w, "home.html", page)
}

func Articles(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Articles"}
	logger.Log("Inside List Articles")
	RenderTemplate(w, "articles.html", page)
}

func Playground(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Playground"}
	logger.Log("Inside Playground")
	RenderTemplate(w, "playground.html", page)
}

func RenderTemplate(w http.ResponseWriter, tmplName string, p interface{}) {
	logger.Log("Rendering template: " + tmplName)
	layout := filepath.Join("tmpl", "layout.html")
	fp := filepath.Join("tmpl", tmplName)

	tmpl := template.Must(template.ParseFiles(fp, layout))

	if err := tmpl.ExecuteTemplate(w, "layout", p); err != nil {
		logger.Log("Couldn't execute template %s : %s", tmpl, err)
		return
	}
}

func main() {
	logger.Log("Starting Server")

	//This is required to serve the css files, or anything we have in /resources
	fs := http.FileServer(http.Dir("./resources"))
	http.Handle("/resources/", http.StripPrefix("/resources/", fs))
	http.HandleFunc("/articles", Articles)
	http.HandleFunc("/playground", Playground)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/", Home)

	logger.Log("Listening on 8000")
	http.ListenAndServe(":8000", nil)

}
