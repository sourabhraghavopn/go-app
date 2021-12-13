package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func (app *App) route() *httprouter.Router {
	router := httprouter.New()
	service := &Service{repo:UrlRepo{conn:app.loadConnection(),logger:app.logger, },}

	router.HandlerFunc(http.MethodPost, "/create", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {http.Error(w, "Error reading request body", http.StatusInternalServerError)}

		request :=CreateShortUrlRequest{}
		json.Unmarshal([]byte(body), &request)

		response, err := json.Marshal(service.createShortUrl(request))
		if err != nil {http.Error(w, "Error reading response body", http.StatusInternalServerError)}

		respondSuccess(w,string(response))
	})
	router.HandlerFunc(http.MethodGet, "/get", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {http.Error(w, "Error reading request body", http.StatusInternalServerError)}

		request :=GetOriginalUrlRequest{}
		json.Unmarshal([]byte(body), &request)

		response, err := json.Marshal(service.getFullUrl(request))
		fmt.Println(string(response))
		if err != nil {http.Error(w, "Error reading response body", http.StatusInternalServerError)}

		respondSuccess(w,string(response))
	})
	return router
}


func respondSuccess(w http.ResponseWriter, input string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(input))
}