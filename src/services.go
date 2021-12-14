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
	dbPort, exist := os.LookupEnv("PORT")
	if !exist || !done {
		service.logger.Print("Properties are missing ")
		return CreateShortUrlResponse{"", "failed"}
	}
	return CreateShortUrlResponse{fmt.Sprintf("http://localhost:%s/%d", dbPort, UrlId), "success"}
}
func (service Service) getFullUrl(request GetOriginalUrlRequest) GetOriginalUrlResponse {
	id, err := extractId(request.ShortUrl)
	url, err := service.repo.get(id)
	if err != nil {
		service.logger.Print("Error found ", err)
		return GetOriginalUrlResponse{"", "failed"}
	}
	return GetOriginalUrlResponse{url.OriginalUrl, "success"}

}
