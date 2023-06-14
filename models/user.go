package models

// User define the structure of the user,
// when necessary info need to be inserted in to database
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
