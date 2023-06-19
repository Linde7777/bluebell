package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
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

// GetPostDetailList return a list of length 2
func GetPostDetailList() (postDetailList []*models.ApiPostDetail, err error) {
	postList, err := mysql.GetPostDetailList()
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
