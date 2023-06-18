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
	postData, err := mysql.GetPostDetailByID(id)
	if err != nil {
		return nil, err
	}

	userData, err := mysql.GetUserByID(postData.AuthorID)
	if err != nil {
		return nil, err
	}

	communityData, rowIsEmpty, err := mysql.GetCommunityDetail(postData.CommunityID)
	if rowIsEmpty && err == nil {
		return nil, err
	}

	apd := new(models.ApiPostDetail)
	apd.CommunityDetailSelected = communityData
	apd.Post = postData
	apd.AuthorName = userData.Username

	return apd, err
}
