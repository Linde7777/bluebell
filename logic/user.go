package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

func SignUp() {
	// 1. is user exist?
	mysql.QueryUserByUsername()
	// 2. generate user id
	snowflake.GenID()
	// 3. DAO, store user in database
	mysql.SignUp()

}
