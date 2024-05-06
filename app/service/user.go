package service

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"
)

func SingUp(req *model.SingUpRequest) error {
	user := model.User{
		ID:       req.UserID,
		Username: req.Username,
		Password: req.Pwd,
	}
	return user.CreateUser()
}

func Login(req *model.LoginRequest) (string, error) {
	u := model.User{}

	err := u.GetUserByID(req.UserID)
	if err != nil {
		return "", err
	}

	if utils.ComparePasswords(u.Password, req.Pwd) {
		return utils.GenerateNewToken(), nil
	}

	return "", pkg.ErrPasswordNotMatch
}
