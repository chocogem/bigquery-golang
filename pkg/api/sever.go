package api

import (
	"github.com/chocogem/bigquery-golang/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type Handlers struct {
	UserHandler *handler.UserHandler
}

func NewServerHTTP(handlers *Handlers) *ServerHTTP {
	engine := gin.New()
	// Healthcheck
	engine.GET("healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})
	api := engine.Group("/user")
	api.GET("/", handlers.UserHandler.FindAll)

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8081")
}
