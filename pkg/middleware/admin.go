package middleware

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AdminOnly(c *gin.Context) {
	adminAccessToken := c.GetHeader("access-token")

	if adminAccessToken != os.Getenv("ADMIN_ACCESS_TOKEN") {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": pkg.ErrOnlyAdminCanAccess.Error(),
		})
		return
	}

	c.Next()
}
