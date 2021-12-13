package main

import (
	"context"
	"github.com/uptrace/bun"
	"log"
)


type App struct {
	logger *log.Logger
	conn   Conn
}
type UrlDetail struct {
	Id          int
	OriginalUrl string
	ShortUrl    int
}
type Conn struct {
	db  *bun.DB
	ctx context.Context
}
type UrlRepo struct {
	logger *log.Logger
	conn   Conn
}
type CreateShortUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}
type CreateShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
	Status      string `json:"status"`
}
type GetOriginalUrlRequest struct {
	ShortUrl string `json:"short_url"`
}
type GetOriginalUrlResponse struct {
	OriginalUrl string `json:"original_url"`
	Status      string `json:"status"`
}
type Service struct {
	repo   UrlRepo
	logger *log.Logger
}
