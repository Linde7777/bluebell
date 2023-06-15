package middlewares

import (
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Assuming Token is stored at the Authorization in Header:
		// Authorization: Bearer: xxx.xxx.xxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "auth is empty",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "wrong auth format",
			})
			c.Abort()
			return
		}

		// parts[1] is tokenString
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}

		c.Set("username", mc.Username)
		c.Next()
	}
}
