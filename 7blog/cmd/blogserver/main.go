package main

import (
	"7blog/cmd/blogserver/router"
	"7blog/internal/conf"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TApp struct {
	ctx      context.Context
	cancel   func()
}

func newApp(logger log.Logger, hs *http.Server) *TApp {
	fmt.Println(hs, logger)
	return new(TApp)
}

func main() {
	logger := log.Logger{}

	var bc conf.Data
	app, cleanup, err := initApp(&bc, logger)
	fmt.Println(app, cleanup, err)
	r := gin.Default()
	router.Register(r)
	r.Run(":8088")
}