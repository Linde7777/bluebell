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

func Login(pl *models.ParamsLogin) (err error) {
	// 1. read from DB, check if user exist,
	// if not, return error, if so, check password,
	// notice that the password is encrypted
	exist, err := mysql.CheckUserExist(pl.Username)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("user does not exist")
	}

	match, err := mysql.CheckPWDMatching(pl)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("password does not match")
	}

	return
}
