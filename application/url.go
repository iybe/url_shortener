package application

import (
	"url_shortener/services"

	"github.com/gin-gonic/gin"
)

type UrlApp struct {
	SizeUrlShortened int
}

func NewUrlApp() *UrlApp {
	return &UrlApp{}
}

type UrlAppInterface interface {
	ShortenUrl(*gin.Context, string) (string, error)
}

var _ UrlAppInterface = &UrlApp{}

func (u *UrlApp) ShortenUrl(c *gin.Context, url string) (string, error) {
	return services.GenerateStringRandom(u.SizeUrlShortened), nil
}
