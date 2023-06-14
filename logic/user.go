package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1. is user exist?
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2. generate user id
	uid := snowflake.GenID()
	u := &models.UserInserted{
		UserID:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	// 3. DAO, store user in database
	if err = mysql.InsertUser(u); err != nil {
		return err
	}

	return
}
