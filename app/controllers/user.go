package controllers

import (
	"errors"
	"fmt"
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/app/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

func SignUp(c *gin.Context) {
	r := new(model.SingUpRequest)

	fmt.Printf("called")

	err := c.Bind(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user_service.SingUp(r)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	}

}

func Login(c *gin.Context) error {
	return errors.New("asdf")
}

func Logout(c *gin.Context) error {
	return errors.New("asdf")
}
