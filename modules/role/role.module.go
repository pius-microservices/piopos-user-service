package role

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRoutesModule(router *gin.Engine, db *gorm.DB, prefix string) {
	roleRepo := NewRepo(db)
	roleService := NewService(roleRepo)
	roleController := NewController(roleService)

	RoleRoutes(router, roleController, prefix)
}
