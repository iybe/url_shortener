package main

import (
	"url_shortener/middleware"

	"github.com/gin-gonic/gin"
)

func routes(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())

	r.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	return r
}
