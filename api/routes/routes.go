package routes

import (
	_ "github.com/pius-microservices/piopos-user-service/docs"
	"github.com/pius-microservices/piopos-user-service/modules/role"
	"github.com/pius-microservices/piopos-user-service/modules/user"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

const (
	APIPrefix = "/api/user-service"
)

func homeHandler(ctx *gin.Context) {
	ctx.JSON(404, gin.H{
		"status":  404,
		"message": "Sorry! Page not found",
	})
}

func RouteApp(router *gin.Engine, db *gorm.DB) error {
	router.GET(APIPrefix+"/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET(APIPrefix, homeHandler)

	role.RoleRoutesModule(router, db, APIPrefix)
	user.UserRoutesModule(router, db, APIPrefix)

	return nil
}
