package routes

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/controllers"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(e *gin.Engine) {
	// Create routes group.
	g := e.Group("/api/v1")

	userRoute := g.Group("/user")
	{
		userRoute.POST("/signup", controllers.SignUp)
		userRoute.POST("/login", controllers.Login)
		userRoute.DELETE("/logout", middleware.ValidatorToken, controllers.Logout)
	}

	accountRoute := g.Group("/account").Use(middleware.ValidatorToken)
	{
		accountRoute.POST("/open", controllers.OpenAccount)
		accountRoute.GET("/:account_id", controllers.GetAccount)
		accountRoute.GET("/", controllers.GetAccountList)
		accountRoute.POST("/deposit", controllers.Deposit)
		accountRoute.POST("/withdraw", controllers.Withdraw)
	}

	portfolioRoute := g.Group("/portfolio").Use(middleware.ValidatorToken)
	{
		portfolioRoute.POST("/request", controllers.CreatePortfolio)
		portfolioRoute.GET("/", controllers.GetPortfolioList)
		portfolioRoute.GET("/:portfolio_id", controllers.GetPortfolio)
	}

	stockRoute := g.Group("/stock")
	{
		stockRoute.POST("/", middleware.AdminOnly, controllers.AddStock)
		stockRoute.GET("/", controllers.GetStockList)
		stockRoute.GET("/:stock_id", controllers.GetStockByCode)
		stockRoute.PATCH("/:stock_id", middleware.AdminOnly, controllers.UpdateStock)
		stockRoute.DELETE("/:stock_id", middleware.AdminOnly, controllers.DeleteStock)
	}
}
