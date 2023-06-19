package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
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
