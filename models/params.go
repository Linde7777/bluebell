package models

// ParamSignUp store the info that is typed by the client
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogin store the info that is needed to be used in Login
type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamsVoteData store the info that is needed to
// be used in voting a post
type ParamsVoteData struct {
	// get UserID from the gin.Context
	PostID string `json:"post_id" binding:"required"`

	// 1 for upvote, -1 for downvote, 0 for not voting
	Direction int8 `json:"direction,string" binding:"required,oneof=1 0 -1"`
}

// ParamsPostList stored the info that
// is needed to get posts
type ParamsPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order" binding:"oneof=time score"`
}

type ParamsCommunityPostList struct {
	ParamsPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
