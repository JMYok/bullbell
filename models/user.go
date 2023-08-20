package models

type User struct {
	UserId   uint64 `db:"user_id"`
	Username string `db:"username"`
	Email    string `db:"password"`
	Gender   int8   `db:"gender"`
	Password string `db:"password"`
}
