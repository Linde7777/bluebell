package controller

import (
	"bluebell/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 1. get argument and validate
	// 2. logic
	logic.SignUp()

	// 3. return
	c.JSON(http.StatusOK, gin.H{"msg": "you have sign up"})
}
