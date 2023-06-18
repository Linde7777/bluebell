package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(pc *models.PostCreated) error {
	pc.ID = snowflake.GenID()
	return mysql.InsertPost(pc)
}
