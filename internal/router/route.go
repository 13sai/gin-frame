package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/13sai/gin-frame/internal/controller"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	g.GET("/", controller.Index)
	g.GET("/healthCheck", controller.HealthCheck)

	return g
}
