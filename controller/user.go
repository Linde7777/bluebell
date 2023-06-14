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
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "invalid parameters",
		})
		return
	}

	// 2. logic
	logic.SignUp()

	// 3. return
	c.JSON(http.StatusOK, gin.H{"msg": "sign up success"})
}
