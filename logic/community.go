package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.CommunityBasicSelected, err error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetailList(id int64) (*models.CommunityDetailSelected, error) {
	data, rowIsEmpty, err := mysql.GetCommunityDetailList(id)
	if rowIsEmpty && err == nil {
		return nil, nil
	}
	return data, err

}
