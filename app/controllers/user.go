package controllers

import (
	"errors"
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/app/service"
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func SignUp(c *gin.Context) {
	r := new(model.SingUpRequest)

	err := c.Bind(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	err = service.SingUp(r)
	if err != nil {
		if errors.Is(err, pkg.ErrDuplicateUser) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "user already exists",
			})
			return
		} else {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}

func Login(c *gin.Context) {
	r := new(model.LoginRequest)

	err := c.Bind(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "request body error"})
	}

	token, err := service.Login(r)

	if err != nil {
		switch {
		case errors.Is(err, pkg.ErrPasswordNotMatch):
			c.JSON(http.StatusBadRequest, gin.H{"message": "password not match"})
			return
		case errors.Is(err, gorm.ErrRecordNotFound):
			c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
	}

	middleware.AddSession(token, r.UserID, c)

	loginLog := model.LoginLog{
		UserID:    r.UserID,
		LoginType: model.Login,
	}

	_ = loginLog.CreateLoginLog()

	c.JSON(http.StatusOK, gin.H{
		"message":      "User logged successfully",
		"access_token": token,
	})
}

func Logout(c *gin.Context) {
	r := c.GetHeader("access-token")

	s := sessions.Default(c)
	uid := s.Get(r).(string)
	s.Set(r, "")
	s.Clear()
	s.Options(sessions.Options{
		MaxAge: -1,
	})
	_ = s.Save()

	loginLog := model.LoginLog{
		UserID:    uid,
		LoginType: model.Logout,
	}

	err := loginLog.CreateLoginLog()
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged successfully",
	})
}
