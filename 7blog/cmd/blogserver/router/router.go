package router

import (
	"7blog/cmd/blogserver/router/handler"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/goslayer/events", (&handler.EventHandler{}).Events)
	engine.POST("goslayer/events/join", (&handler.EventHandler{}).JoinAEvent)
}