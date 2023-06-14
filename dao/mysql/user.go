package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func QueryUserByID() {

}

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user is already exists")
	}

	return

}

func QueryUserByUsername() {

}

// InsertUser will encrypt the models.UserInserted.Password
func InsertUser(u *models.UserInserted) error {
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
