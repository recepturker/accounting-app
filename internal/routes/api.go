package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/recepturker/accounting-app/config"
	"github.com/recepturker/accounting-app/internal/handlers"
	"github.com/recepturker/accounting-app/internal/repositories"
	"github.com/recepturker/accounting-app/internal/services"
)

func RegisterRoutes(router *gin.Engine) {
	config.ConnectDB()
	var userRepo *repositories.UserRepository = repositories.NewUserRepository(config.DB)
	var userService *services.UserService = services.NewUserService(userRepo)
	var userHandler *handlers.UserHandler = handlers.NewUserHandler(userService)
	var authHandler *handlers.AuthHandler = handlers.NewAuthHandler(userService)

	var api *gin.RouterGroup = router.Group("/api")
	{
		api.POST("/register", userHandler.RegisterUser)
		api.POST("/login", authHandler.Login)
		// api.GET("/users", userHandler.GetUsers)
		// api.GET("/users/:id", userHandler.GetUser)
	}
}
