package moudel


type User struct {
	username string `json:"username" binding:"required"`
	password string `json:"password" binding:"required"`
}
