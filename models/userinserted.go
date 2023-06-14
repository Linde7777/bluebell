package models

// UserInserted store the info that is needed
// to be inserted into database
type UserInserted struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
