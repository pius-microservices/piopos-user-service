package user

import (
	// "github.com/pius-microservices/piopos-user-service/config"
	"github.com/pius-microservices/piopos-user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *userController, prefix string) {
	// envCfg := config.LoadConfig()

	userGroup := router.Group(prefix + "/user")
	{
		userGroup.POST("/signup", func(ctx *gin.Context) {
			controller.SignUp(ctx)
		})
		userGroup.PUT("/verify-account", func(ctx *gin.Context) {
			controller.VerifyAccount(ctx)
		})
		userGroup.PUT("/send-otp", func(ctx *gin.Context) {
			controller.SendNewOTPCode(ctx)
		})
		userGroup.PUT("/update", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.UpdateUserProfile(ctx)
		})
		userGroup.PUT("/update-password", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.UpdatePassword(ctx)
		})

		userGroup.GET("/", func(ctx *gin.Context) {
			controller.GetUsers(ctx)
		})
		userGroup.GET("/:id", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetUserById(ctx)
		})

		userGroup.GET("/profile", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.GetProfile(ctx)
		})
		userGroup.GET("/get-user-by-email", func(ctx *gin.Context) {
			controller.GetUserByEmail(ctx)
		})
	}
}
