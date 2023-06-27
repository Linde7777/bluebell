package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrUserNotLogin = errors.New("user has not login")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrUserNotLogin
		return
	}

	return
}
