package model

type Admin struct {
	AdminID       int    `json:"Admin_id" form:"Admin_id"`
	AdminName     string `json:"Admin_name" form:"Admin_name"`
	AdminPassword string `json:"Admin_password" form:"Admin_password"`
	AdminLevel    int    `json:"Admin_level" form:"Admin_level"`
}

type PostAdmin struct {
	PostId string  `json:"post_id"`
	Admin  []Admin `json:"data"`
}
