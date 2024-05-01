package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// SwaggerRoute func for describe group of API Docs routes.
func SwaggerRoute(e *gin.Engine) {
	// Create routes group.
	e.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/swagger/index.html")
	})
	// Routes for GET method:
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // get one user by ID
}
