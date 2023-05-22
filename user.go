package Smarthouse_server

type User struct {
	Id            int    `json:"-" db:"id"`
	Name          string `json:"name" binding:"required"`
	Username      string `json:"username" binding:"required"`
	Password_hash string `json:"password_hash" binding:"required"`
}
