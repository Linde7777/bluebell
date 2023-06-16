package middlewares

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"github.com/gin-gonic/gin"
	"strings"
)

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Assuming Token is stored at the Authorization in Header:
		// Authorization: Bearer xxx.xxx.xxx
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1] is tokenString
		mc, err := jwt.ParseToken(parts[1])
		//todo dealing with refreshToken
		// todo specify a code for accessToken expire
		if jwt.IsTimeExpireErr(err) {
			controller.ResponseErrorWithMsg(c, controller.CodeInvalidToken,
				"AccessToken/RefreshToken has expired")
			c.Abort()
			return
		}
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next()
	}
}
