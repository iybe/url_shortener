package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r = routes(r)

	r.Run(":8888")
}
