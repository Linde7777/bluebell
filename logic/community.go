package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.CommunityBasicSelected, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetailSelected, error) {
	data, rowIsEmpty, err := mysql.GetCommunityDetail(id)
	if rowIsEmpty && err == nil {
		return nil, nil
	}
	return data, err

}
