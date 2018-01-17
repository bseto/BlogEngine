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
	Title      string   `json:"title"`
	Path       string   `json:"path"`
	Tags       []string `json:"tags, omitempty"`
	Body       string   `json:"body, omitempty"`
	CreateDate string   `json:"create_date"`
}

type ListYMLStruct struct {
	List []YMLStruct `yaml:"articles"`
}
type YMLStruct struct {
	Title      string   `yaml:"title"`
	Path       string   `yaml:"path"`
	CreateDate string   `yaml:"create_date"`
	Tags       []string `yaml:"tags"`
}

func Home(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Home"}
	logger.Log("Inside Home")
	RenderTemplate(w, page,
		filepath.Join("tmpl", "home.html"))
}

func Articles(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Articles"}
	logger.Log("Inside List Articles")
	RenderTemplate(w, page,
		filepath.Join("tmpl", "articles.html"))
}

func Playground(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Playground"}
	logger.Log("Inside Playground")
	RenderTemplate(w, page,
		filepath.Join("tmpl", "playground.html"))
}

func RenderTemplate(w http.ResponseWriter, p interface{}, tmplName ...string) {
	logger.Log("Rendering templates:%v ", tmplName)
	layout := filepath.Join("tmpl", "layout.html")
	tmplName = append(tmplName, layout)

	tmpl, err := template.ParseFiles(tmplName...)
	if err != nil {
		//If err, then we will make tmpl 404
		notFound := filepath.Join("tmpl", "404.html")
		tmpl = template.Must(template.ParseFiles(layout, notFound))
	}

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
