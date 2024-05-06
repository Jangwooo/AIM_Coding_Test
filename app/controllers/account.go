package controllers

import (
	"errors"
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/app/service"
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func OpenAccount(c *gin.Context) {
	uid := c.GetHeader("user_id")

	err := service.OpenAccount(uid)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "account opened",
	})
}

func GetAccountList(c *gin.Context) {
	uid := c.GetHeader("user_id")

	as, err := service.GetAccountList(uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": pkg.ErrRecordNotFound.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    as,
	})
}

func GetAccount(c *gin.Context) {
	uid := c.GetHeader("user_id")
	aid := c.Param("account_id")

	a, err := service.GetAccountInfo(aid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": pkg.ErrRecordNotFound.Error(),
			})
			return
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			return
		}
	}

	if a.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    a,
	})
}

func Deposit(c *gin.Context) {
	uid := c.GetHeader("user_id")
	req := new(model.DepositRequest)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	a, err := service.GetAccountInfo(req.AccountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account not found",
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

	if a.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
		return
	}

	err = service.Deposit(uid, req.AccountID, req.Amount)
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit success",
	})

}

func Withdraw(c *gin.Context) {
	uid := c.GetHeader("user_id")
	req := new(model.WithdrawRequest)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	a, err := service.GetAccountInfo(req.AccountID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account not found",
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

	if a.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "forbidden",
		})
		return
	}

	err = service.Withdraw(uid, req.AccountID, req.Amount)
	if err != nil {
		if errors.Is(err, pkg.ErrBalanceNotEnough) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": pkg.ErrBalanceNotEnough.Error(),
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	a = &model.Account{
		ID: req.AccountID,
	}

	a, _ = a.GetAccountByID()

	t := model.Transaction{
		AccountID:    req.AccountID,
		Amount:       req.Amount,
		Type:         model.TransactionTypeWithdrawal,
		AfterBalance: a.Balance,
	}

	err = t.CreateTransaction()
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "withdraw success",
	})
}
