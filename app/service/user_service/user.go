package user_service

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
)

func SingUp(req *model.SingUpRequest) error {
	user := model.User{}
	err := user.CreateUser(req.Id, req.Pwd, req.Username)
	if err != nil {
		return err
	}
	return nil
}

func Login(req *model.LoginRequest) (string, error) {
	return "", nil
}

func Logout(req *model.LogoutRequest) error {
	return nil
}
