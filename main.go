package main

import (
	"fmt"
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/Jangwooo/AIM_Coding_Test/pkg/routes"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"

	_ "github.com/Jangwooo/AIM_Coding_Test/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .app.env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	_ = godotenv.Load(".app.env")
	db := database.OpenDBConnection()

	err := db.AutoMigrate(
		&model.User{},
		&model.Account{},
		&model.Transaction{},
		&model.Portfolio{},
		&model.PortfolioItem{},
		&model.Stock{},
		&model.LoginLog{})
	if err != nil {
		log.Fatal(fmt.Errorf("error initializing database: %w", err))
	}

	log.Print("Connected to database & Successfully initialized")

	// Define Fiber config.

	// Define a new Gin app with config.
	app := gin.Default()

	// Middlewares.

	// Routes.
	routes.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app) // Register a public routes for app.

	// Start server (with or without graceful shutdown).
	switch os.Getenv("STAGE_STATUS") {
	case "dev":
		gin.SetMode("debug")
		utils.StartServer(app)

	case "prod":
		gin.SetMode("release")
		utils.StartServerWithGracefulShutdown(app)

	}
}
