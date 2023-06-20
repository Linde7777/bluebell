package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"errors"
)

var (
	ErrUserExist    = errors.New("user is already exist")
	ErrUserNotExist = errors.New("user does not exist")
	ErrPwdNotMatch  = errors.New("password does not match")
)

func SignUp(p *models.ParamSignUp) (err error) {
	exist, err := mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}
	if exist {
		return ErrUserExist
	}

	uid := snowflake.GenID()
	u := &models.User{
		UserID:   uid,
		Username: p.Username,
		Password: p.Password,
	}

	if err = mysql.InsertUser(u); err != nil {
		return err
	}

	return
}

func Login(pl *models.ParamsLogin) (user *models.User, err error) {
	// 1. read from DB, check if user exist,
	// if not, return error, if so, check password,
	// notice that the password is encrypted
	exist, err := mysql.CheckUserExist(pl.Username)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, ErrUserNotExist
	}

	match, err := mysql.CheckPWDMatching(pl)
	if err != nil {
		return nil, err
	}
	if !match {
		return nil, ErrPwdNotMatch
	}

	userID, err := mysql.GetUserIDByName(pl.Username)
	if err != nil {
		return nil, err
	}

	accToken, refToken, err := jwt.GenToken(userID, pl.Username)
	if err != nil {
		return nil, err
	}

	user = new(models.User)
	user.UserID = userID
	user.Username = pl.Username
	user.AccessToken = accToken
	user.RefreshToken = refToken
	return
}
