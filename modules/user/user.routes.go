package user

import (
	"github.com/pius-microservices/piopos-user-service/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, controller *userController, prefix string) {

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
		userGroup.POST("/refresh-token", func(ctx *gin.Context) {
			controller.CreateRefreshToken(ctx)
		})
		userGroup.DELETE("/refresh-token", func(ctx *gin.Context) {
			controller.DeleteRefreshToken(ctx)
		})
		userGroup.POST("/validate-refresh-token", func(ctx *gin.Context) {
			controller.ValidateRefreshToken(ctx)
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
