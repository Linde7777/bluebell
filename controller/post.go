package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	// 1. get parameter
	pc := new(models.Post)
	if err := c.ShouldBindJSON(pc); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeNeedLogin, err.Error())
		return
	}
	pc.AuthorID = userID

	if err := logic.CreatePost(pc); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	// get params
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandler: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID: ", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}

	ResponseSuccess(c, data)
}
