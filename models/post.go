package models

import "time"

// Post store the info that is needed
// to create a post
type Post struct {
	// JavaScript can not store int64, need to convert ID to string
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail provide the api for frontend,
// store the detail info of a post
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VoteCount  int64  `json:"vote_count"`
	*Post
	*CommunityDetailSelected
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
