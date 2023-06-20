package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
)

func CheckUserExist(username string) (exist bool, err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return false, err
	}

	return count > 0, nil
}

// InsertUser will encrypt the models.User.Password
func InsertUser(u *models.User) error {
	u.Password = encrypt(u.Password)
	// 1. run sql
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err := db.Exec(sqlStr, u.UserID, u.Username, u.Password)
	return err
}

const secret = "longpeng96@gmail.com"

func encrypt(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func CheckPWDMatching(pl *models.ParamsLogin) (match bool, err error) {
	sqlStr := "select password from user where username=?"
	var password string
	if err = db.Get(&password, sqlStr, pl.Username); err != nil {
		return false, err
	}

	return encrypt(pl.Password) == password, nil
}

func GetUserIDByName(username string) (userID int64, err error) {
	sqlStr := "select user_id from user where username=?"
	if err = db.Get(&userID, sqlStr, username); err != nil {
		return -1, err
	}
	return userID, nil
}

func GetUserByID(userID int64) (*models.User, error) {
	u := new(models.User)
	sqlStr := "select user_id,username from user where user_id=?"
	err := db.Get(u, sqlStr, userID)
	return u, err
}
