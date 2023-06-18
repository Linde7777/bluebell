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

func GetPostDetailByID(id int64) (*models.Post, error) {
	return mysql.GetPostDetailByID(id)
}
