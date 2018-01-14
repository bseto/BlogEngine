package main

import (
	"encoding/json"
	"github.com/bseto/BlogEngine/logger"
	"net/http"
)

func ListArticles(w http.ResponseWriter, req *http.Request) {
	logger.Log("List Articles was Called")
	article := Article{"TestTitle", "Some Fake Lorem Ipsum", "2018-01-14"}
	json, err := json.Marshal(article)
	if err != nil {
		logger.Error(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	logger.Log(string(json))
}
