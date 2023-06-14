package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() (engine *gin.Engine) {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello there")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong!")
	})
	r.POST("/signup", controller.SignUpHandler)
	return r
}
