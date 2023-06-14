package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		msg := "SignUp failed to bind JSON"
		zap.L().Error(msg, zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("fail to signup: " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "sign up success"})
}

func LoginHandler(c *gin.Context) {
	ul := new(models.UserLogin)
	if err := c.ShouldBindJSON(ul); err != nil {
		msg := "LoginHandler fail to bind JSON"
		zap.L().Error(msg, zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if err := logic.Login(ul); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": fmt.Sprintf("fail to login: " + err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, "login success")
}
