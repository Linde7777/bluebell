package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. get parameters and validate them
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		msg := "SignUp failed to bind JSON"
		zap.L().Error(msg, zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}

	if len(p.Password) == 0 || len(p.RePassword) == 0 || len(p.Username) == 0 {
		msg := "SignUp with empty parameters"
		zap.L().Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}

	if p.Password != p.RePassword {
		msg := "Password does not match RePassword"
		zap.L().Error(msg)
		c.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
	}

	// 2. logic
	logic.SignUp(p)

	// 3. return
	c.JSON(http.StatusOK, gin.H{"msg": "sign up success"})
}
