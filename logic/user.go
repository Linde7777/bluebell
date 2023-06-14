package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1. is user exist?
	mysql.QueryUserByUsername()
	// 2. generate user id
	snowflake.GenID()
	// 3. DAO, store user in database
	mysql.SignUp()

}
