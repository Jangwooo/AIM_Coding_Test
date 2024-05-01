package routes

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/controllers"
	"github.com/gin-gonic/gin"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(e *gin.Engine) {
	// Create routes group.
	g := e.Group("/api/v1")

	userRoute := g.Group("/user")
	{
		userRoute.POST("/signup", controllers.SignUp)
		userRoute.POST("/login")
		userRoute.DELETE("/logout")
	}

	accountRoute := g.Group("/account")
	{
		accountRoute.POST("/open")
		accountRoute.GET("/balance/:account_id")
		accountRoute.POST("/deposit")
		accountRoute.POST("/withdraw")
	}

	portfolioRoute := g.Group("/portfolio")
	{
		portfolioRoute.POST("/request")
		portfolioRoute.GET("/:user_id")
	}

	stockRoute := g.Group("/stock")
	{
		stockRoute.POST("/")
		stockRoute.GET("/:stock_code")
		stockRoute.PATCH("/update_price")
		stockRoute.DELETE("/:stock_code")
	}
}
