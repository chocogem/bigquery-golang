package api

import (
	"github.com/chocogem/bigquery-golang/pkg/api/handler"
	"github.com/chocogem/bigquery-golang/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

type Handlers struct {
	UserHandler *handler.UserHandler
}
type Middlewares struct {
	AuthenticationHandler *middleware.AuthenticationHandler
	ErrorHandler   *middleware.ErrorHandler
}

func NewServerHTTP(handlers *Handlers, middlewares *Middlewares) *ServerHTTP {
	engine := gin.New()
	// Healthcheck
	engine.GET("healthcheck", func(c *gin.Context) {
		c.String(200, "OK")
	})
	engine.Use(middlewares.ErrorHandler.Handler())
	//Admin endpoint
	// Auth middleware
	// adminGroup := engine.Group("/admin", middlewares.AuthenticationHandler.Authentication())
	// adminGroup.GET("/user", handlers.UserHandler.FindAll)

	api := engine.Group("/user")
	api.GET("/", handlers.UserHandler.FindAll)

	return &ServerHTTP{engine: engine}
}
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8081")
}
