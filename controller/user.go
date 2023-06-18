package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	ps := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(ps); err != nil {
		zap.L().Error("SignUPHandler: ", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	if err := logic.SignUp(ps); err != nil {
		zap.L().Error("SignUPHandler: ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, "sign up success")
}

func LoginHandler(c *gin.Context) {
	pl := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(pl); err != nil {
		zap.L().Error("LoginHandler: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	accToken, refToken, err := logic.Login(pl)
	if err != nil {
		zap.L().Error("logic.Login: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, fmt.Sprintf("AccessToken: "+accToken, "\nRefreshToken: "+refToken))
}
