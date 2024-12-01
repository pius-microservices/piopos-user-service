package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutesModule(router *gin.Engine, db *gorm.DB, prefix string)  {
	userRepo := NewRepo(db)
	userService := NewService(userRepo)
	userController := NewController(userService)

	UserRoutes(router, userController, prefix)
}