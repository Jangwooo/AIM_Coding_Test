package middleware

import (
	"github.com/Jangwooo/AIM_Coding_Test/pkg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ValidatorToken(c *gin.Context) {
	at := c.GetHeader("access-token")
	session := sessions.Default(c)

	v, ok := session.Get(at).(string)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": pkg.ErrSessionExpired.Error(),
		})
		return
	}

	AddSession(at, v, c)

	c.Request.Header.Add("user_id", v)

	c.Next()
}

func AddSession(k string, v interface{}, c *gin.Context) {
	session := sessions.Default(c)

	session.Set(k, v)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24 * 7,
	})
	_ = session.Save()
}
