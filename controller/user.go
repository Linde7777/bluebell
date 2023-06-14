package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	ps := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(ps); err != nil {
		msg := "SignUp failed to bind JSON"
		zap.L().Error(msg, zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	if err := logic.SignUp(ps); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	ResponseSuccess(c, "sign up success")
}

func LoginHandler(c *gin.Context) {
	pl := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(pl); err != nil {
		msg := "LoginHandler fail to bind JSON"
		zap.L().Error(msg, zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	if err := logic.Login(pl); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	ResponseSuccess(c, "login success")
	c.JSON(http.StatusOK, "login success")
}
