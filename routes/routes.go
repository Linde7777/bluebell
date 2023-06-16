package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUp(mode string) (engine *gin.Engine) {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello there")
	})
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})

	return r
}
