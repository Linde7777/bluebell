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
			"msg": err.Error(),
		})
		return
	}

	// 2. logic
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "fail to signup",
		})
		return
	}

	// 3. return
	c.JSON(http.StatusOK, gin.H{"msg": "sign up success"})
}
