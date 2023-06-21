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
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// GetPostDetailListHandler will return a list of posts details,
// if param "page" is empty, the page will be set to 1
// if param size is empty, the size will be set to 2
func GetPostDetailListHandler(c *gin.Context) {
	targetPageNumberStr := c.Query("page")
	pageSizeStr := c.Query("size")
	if targetPageNumberStr == "" {
		targetPageNumberStr = "1"
	}
	if pageSizeStr == "" {
		pageSizeStr = "2"
	}

	var targetPageNumber int64
	targetPageNumber, err := strconv.ParseInt(targetPageNumberStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	var pageSize int64
	pageSize, err = strconv.ParseInt(pageSizeStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostDetailList(targetPageNumber, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostDetailList: ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

func PostVoteController(c *gin.Context) {
	p := new(models.ParamsVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ShouldBindJSON: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
