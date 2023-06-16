package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.CommunityBasicSelected, err error) {
	return mysql.GetCommunityList()
}
