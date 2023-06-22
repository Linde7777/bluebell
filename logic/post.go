package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
	"strconv"
)

func CreatePost(pc *models.Post) error {
	pc.ID = snowflake.GenID()
	return mysql.InsertPost(pc)
}

func GetPostDetailByID(id int64) (*models.ApiPostDetail, error) {
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		return nil, err
	}

	userData, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		return nil, err
	}

	communityData, rowIsEmpty, err := mysql.GetCommunityDetail(post.CommunityID)
	if rowIsEmpty && err == nil {
		return nil, err
	}

	apd := new(models.ApiPostDetail)
	apd.CommunityDetailSelected = communityData
	apd.Post = post
	apd.AuthorName = userData.Username

	return apd, err
}

// GetPostDetailList return a list of posts,
// start at targetPageNumber, with length pageSize
func GetPostDetailList(targetPageNumber, pageSize int64) (postDetailList []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostDetailList(targetPageNumber, pageSize)
	if err != nil {
		return nil, err
	}

	postDetailList = make([]*models.ApiPostDetail, 0, len(postList))

	for _, post := range postList {
		userData, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID: ", zap.Error(err))
			continue
		}

		communityData, rowIsEmpty, err := mysql.GetCommunityDetail(post.CommunityID)
		if rowIsEmpty && err == nil {
			zap.L().Error("mysql.GetCommunityDetail: ", zap.Error(err))
			continue
		}

		apd := new(models.ApiPostDetail)
		apd.CommunityDetailSelected = communityData
		apd.Post = post
		apd.AuthorName = userData.Username

		postDetailList = append(postDetailList, apd)
	}

	return
}

/*
When a user upvotes a post (direction is 1):
The user hasn't voted before and is now upvoting the post.
The user previously downvoted the post and is now upvoting it instead.

When a user doesn't vote (direction is 0):
The user previously upvoted the post and is now cancelling their vote.
The user previously downvoted the post and is now cancelling their vote.

When a user downvotes a post (direction is -1):
The user hasn't voted before and is now downvoting the post.
The user previously upvoted the post and is now downvoting it instead.

Voting restrictions:
Users are allowed to vote on each post within one week of its publication.
After one week, no further votes are allowed.

After the deadline, move the data of voting in redis to mysql,
and delete the KeyPostVotedPrefix key in redis
*/

func VoteForPost(userID int64, p *models.ParamsVoteData) error {
	zap.L().Debug("VoteForPost: ",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)),
		p.PostID, float64(p.Direction))
}
