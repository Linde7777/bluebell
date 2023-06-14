package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"errors"
)

func SignUp(p *models.ParamSignUp) (err error) {
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("user is already exist")
	}

	uid := snowflake.GenID()
	u := &models.UserInserted{
		UserID:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.InsertUser(u); err != nil {
		return err
	}

	return
}
