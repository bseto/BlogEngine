package main

import (
	"github.com/bseto/BlogEngine/logger"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
)

type GeneralPage struct {
	ActiveTab string
}

type Article struct {
	Title      string `json:"title"`
	Body       string `json:"body, omitempty"`
	CreateDate string `json:"create_date"`
}

type ListYMLStruct struct {
	List []YMLStruct `yaml:"articles"`
}
type YMLStruct struct {
	Title      string `yaml:"title"`
	Path       string `yaml:"path"`
	CreateDate string `yaml:"create_date"`
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

	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/articles", Articles)
	r.HandleFunc("/playground", Playground)
	r.HandleFunc("/home", Home)
	r.HandleFunc("/article/{article-title}", GetArticle)
	r.HandleFunc("/api/list_articles", ListArticles)

	http.Handle("/", r)

	logger.Log("Listening on 8000")
	http.ListenAndServe(":8000", nil)

}
