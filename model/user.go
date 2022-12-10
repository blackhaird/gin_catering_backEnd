package model

type User struct {
	UserID       int    `json:"user_id" form:"user_id"`
	UserName     string `json:"user_name" form:"user_name"`
	UserPassword string `json:"user_password" form:"user_password"`
	UserLevel    int    `json:"user_level" form:"user_level"`
}

type PostUser struct {
	PostId string `json:"post_id"`
	User   []User `json:"data"`
}
