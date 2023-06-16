package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CommunityHandler get all community,
// return <community_id, community_name> pairs
func CommunityHandler(c *gin.Context) {
	communityList, err := logic.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	communityDetail, err := logic.GetCommunityDetail(id)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, communityDetail)
}
