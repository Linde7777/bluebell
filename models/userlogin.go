package models

// UserLogin store the info that is needed to be login
type UserLogin struct {
	Username string `db:"username"`
	Password string `db:"password"`
}
