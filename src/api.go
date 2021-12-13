package main

import (
	"fmt"
	"os"
)

func (service Service) createShortUrl(request CreateShortUrlRequest) CreateShortUrlResponse {
	UrlId := service.repo.getLatestSequence()
	done := service.repo.insert(UrlDetail{
		Id:          UrlId,
		OriginalUrl: string(request.OriginalUrl),
		ShortUrl:    UrlId,
	})
	dbPort, exist := os.LookupEnv("DB_PORT")
	if !exist || !done {
		service.logger.Print("properties are missing")
		return CreateShortUrlResponse{"", "failed"}
	}
	return CreateShortUrlResponse{
		fmt.Sprintf("http://localhost:%s/%d", dbPort, UrlId),
		"success"}
}
func (service Service) getFullUrl(request GetOriginalUrlRequest) GetOriginalUrlResponse {
	id, err := extractId(request.ShortUrl)
	url, err := service.repo.get(id)
	if err != nil {
		return GetOriginalUrlResponse{"", "failed"}
	} else {
		fmt.Println("No error found ")
		return GetOriginalUrlResponse{url.OriginalUrl, "success"}
	}
}
