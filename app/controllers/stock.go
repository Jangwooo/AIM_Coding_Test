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

func AddStock(c *gin.Context) {
	r := new(model.CreateStockRequest)
	err := c.Bind(r)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	err = service.AddStock(r)
	if err != nil {
		if errors.Is(err, pkg.ErrDuplicateStockID) {
			c.JSON(http.StatusConflict, gin.H{
				"message": "stock already exists",
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
		"message": "stock created",
	})
}

func GetStockByCode(c *gin.Context) {
	sid := c.Param("stock_id")

	s, err := service.GetStockByCode(sid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "stock not found",
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
		"data":    s,
	})
}

func UpdateStock(c *gin.Context) {
	r := new(model.UpdateStockRequest)
	sid := c.Param("stock_id")
	err := c.Bind(r)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "request body error",
		})
		return
	}

	_, err = service.GetStockByCode(sid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": pkg.ErrRecordNotFound.Error(),
			})
			return
		}
	}

	err = service.UpdateStock(r, sid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "stock not found",
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
		"message": "stock updated",
	})
}

func DeleteStock(c *gin.Context) {
	sid := c.Param("stock_id")

	_, err := service.GetStockByCode(sid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": pkg.ErrRecordNotFound.Error(),
			})
			return
		} else {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
		}
	}

	err = service.DeleteStock(sid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "stock not found",
			})
			return
		}
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "stock deleted",
	})
}

func GetStockList(c *gin.Context) {
	ss, err := service.GetStockList()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": pkg.ErrRecordNotFound.Error(),
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
		"data":    ss,
	})
}
