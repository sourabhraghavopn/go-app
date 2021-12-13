package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
func extractId(url string) (int, error) {
	list := strings.Split(url, "/")
	if len(list) > 3 {
		id_int, _ := strconv.Atoi(list[3])
		return id_int, nil
	}
	return 0, errors.New("incorrect url")
}
