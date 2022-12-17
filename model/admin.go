package model

type Admin struct {
	AdminNo       int    `json:"admin_no"`
	UserID        string `json:"user_id" `
	AdminPassword int    `json:"admin_password" `
}

type PostAdmin struct {
	PostId string  `json:"post_id"`
	Admin  []Admin `json:"data"`
}
