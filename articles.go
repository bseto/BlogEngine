package main

import (
	"encoding/json"
	"github.com/bseto/BlogEngine/logger"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"path"
)

func readYML() (ListYMLStruct, error) {
	ymlFile, err := ioutil.ReadFile("./articles/articles.yml")

	var listStruct ListYMLStruct
	err = yaml.Unmarshal(ymlFile, &listStruct)
	if err != nil {
		logger.Error("Err: %v\n", err)
		return listStruct, err
	}

	logger.Log("ymlStruct is: %v", listStruct)
	return listStruct, nil
}

func ListArticles(w http.ResponseWriter, req *http.Request) {
	articlesYML, err := readYML()
	if err != nil {
		logger.Error("Err: %v\n", err)
		return
	}
	var articles []Article
	for _, articleDetails := range articlesYML.List {
		tempArticle := Article{Title: articleDetails.Title,
			CreateDate: articleDetails.CreateDate,
			Path:       articleDetails.Path,
			Tags:       articleDetails.Tags}

		articles = append(articles, tempArticle)
	}
	json, err := json.Marshal(articles)
	logger.Log("sending json: %v\n", articles)
	if err != nil {
		logger.Error("Err: %v\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func GetArticle(w http.ResponseWriter, req *http.Request) {
	page := GeneralPage{ActiveTab: "Articles"}
	vars := mux.Vars(req)
	logger.Log("URL was" + vars["article-title"])
	if val, ok := vars["article-title"]; ok {
		_, err := ioutil.ReadFile(path.Join("/articles/", val+".html"))
		if err != nil {
			logger.Error("Err: %v\n", err)
			RenderTemplate(w, "404.html", page)
			return
		}
	} else {
		logger.Error("There was no article-title provided")
		RenderTemplate(w, "404.html", page)
	}
}
