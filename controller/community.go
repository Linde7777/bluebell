package controller

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
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
