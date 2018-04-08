package models

// User is an application user
type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
}
