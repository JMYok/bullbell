package models

type User struct {
	UserId   uint64 `db:"user_id"`
	Username string `db:"username"`
	Email    string `db:"email"`
	Gender   int8   `db:"gender"`
	Password string `db:"password"`
}
