package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(c *gin.Context) {
	// 1. get parameter
	pc := new(models.PostBasic)
	if err := c.ShouldBindJSON(pc); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
	}

	userID, err := getCurrentUser(c)
	if err != nil {
		ResponseErrorWithMsg(c, CodeNeedLogin, err.Error())
		return
	}
	pc.AuthorID = userID

	if err = logic.CreatePost(pc); err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, nil)
}

func GetPostDetailHandler(c *gin.Context) {
	id := c.Param("id")
	data, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID: ", zap.Error(err))
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

	userID, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("getCurrentUser: ", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}

	if err = logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, nil)
}

// GetPostDetailListHandler return posts list
// by time or by scores.
// If community_id is not empty, it will return the
// posts under the community, otherwise it will
// return all the posts.
// Example Usage:
// api/v1/posts2?page=1&size=2&order=time&community_id=123456
// api/v1/posts2?page=1&size=2&order=scores&community_id=0
func GetPostDetailListHandler(c *gin.Context) {
	p := &models.ParamsPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("c.ShouldBindQuery: ", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	postList, err := logic.GetPostDetailList(p)
	if err != nil {
		zap.L().Error("logic.GetAllPostDetail: ", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)
}
