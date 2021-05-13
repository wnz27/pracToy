package main

import (
	"7blog/cmd/blogserver/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Register(r)
	r.Run(":8080")
}