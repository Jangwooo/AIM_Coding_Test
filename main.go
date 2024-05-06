package main

import (
	"github.com/Jangwooo/AIM_Coding_Test/app/model"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/middleware"
	"github.com/Jangwooo/AIM_Coding_Test/platform/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/Jangwooo/AIM_Coding_Test/pkg/routes"
	"github.com/Jangwooo/AIM_Coding_Test/pkg/utils"

	_ "github.com/Jangwooo/AIM_Coding_Test/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .app.env file automatically
)

// @title AIM 코딩 테스트 과제 API 문서
// @version 1.0
// @description 코딩 테스트의 원활한 채점을 위한 API 문서 입니다
// @BasePath /api/v1

func main() {
	err := godotenv.Load(".app.env")
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Account{},
		&model.Transaction{},
		&model.Portfolio{},
		&model.PortfolioItem{},
		&model.Stock{},
		&model.LoginLog{},
		&model.StockTransaction{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Connected to database & Successfully initialized")

	// Define Fiber config.

	// Define a new Gin app with config.
	app := gin.Default()

	// Middlewares.
	app.Use(middleware.Timeout())

	store, _ := redis.NewStore(100, "tcp", utils.ConnectionURLBuilder("redis"), os.Getenv("REDIS_PASSWORD"), []byte("secret"))
	app.Use(sessions.Sessions("access", store))

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
