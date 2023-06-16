package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.CommunitySelected, err error) {
	return mysql.GetCommunityList()
}
