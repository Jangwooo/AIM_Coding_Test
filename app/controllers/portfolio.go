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

func CreatePortfolio(c *gin.Context) {
	uid := c.GetHeader("user_id")
	req := new(model.CreatePortfolioRequest)

	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	a := &model.Account{
		ID: req.AccountID,
	}

	a, err = a.GetAccountByID()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account not found",
			})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	if a.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": pkg.ErrForbiddenResourceAccess.Error(),
		})
		return
	}

	err = service.CreatePortfolio(uid, req.AccountID, req.RiskType, a)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "portfolio created",
	})
}

func GetPortfolio(c *gin.Context) {
	pid := c.Param("portfolio_id")
	uid := c.GetHeader("user_id")

	p := &model.Portfolio{
		ID: pid,
	}

	p, err := p.GetPortfolioByID()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "portfolio not found",
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

	if p.UserID != uid {
		c.JSON(http.StatusForbidden, gin.H{
			"message": pkg.ErrForbiddenResourceAccess.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    p,
	})
}

func GetPortfolioList(c *gin.Context) {
	uid := c.GetHeader("user_id")

	p := &model.Portfolio{
		UserID: uid,
	}

	plist, err := p.GetPortfolioList()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "portfolio not found",
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

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    plist,
	})
}
