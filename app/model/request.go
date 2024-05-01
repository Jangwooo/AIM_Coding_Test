package model

type SingUpRequest struct {
	Id       string `json:"id" binding:"required"`
	Pwd      string `json:"pwd" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Id  string `json:"id" binding:"required"`
	Pwd string `json:"pwd" binding:"required"`
}

type LogoutRequest struct {
}
