package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func (app *App) route(repo UrlRepo) *httprouter.Router {
	router := httprouter.New()
	service := &Service{repo: repo,logger:app.logger}

	router.HandlerFunc(http.MethodPost, "/create", func(w http.ResponseWriter, r *http.Request) {
		request := extractCreateRequest(w, r)
		sendResponse(w, service.createShortUrl(request))
	})
	router.HandlerFunc(http.MethodPost, "/get", func(w http.ResponseWriter, r *http.Request) {
		request := extractGetRequest(w, r)
		sendResponse(w, service.getFullUrl(request))
	})
	return router
}

func extractCreateRequest(w http.ResponseWriter, r *http.Request) CreateShortUrlRequest {
	request := CreateShortUrlRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	json.Unmarshal([]byte(body), &request)
	return request
}

func extractGetRequest(w http.ResponseWriter, r *http.Request) GetOriginalUrlRequest {
	request := GetOriginalUrlRequest{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}
	json.Unmarshal([]byte(body), &request)
	return request
}

func sendResponse(w http.ResponseWriter, response interface{}) {
	responseString, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseString))
}
