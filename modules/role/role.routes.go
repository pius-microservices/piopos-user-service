package role

import "github.com/gin-gonic/gin"

func RoleRoutes(router *gin.Engine, controller *roleController, prefix string) {
	roleGroup := router.Group(prefix + "/role")
	{
		roleGroup.POST("", func(ctx *gin.Context) {
			controller.AddRole(ctx)
		})
		roleGroup.GET("", func(ctx *gin.Context) {
			controller.GetRoles(ctx)
		})
		roleGroup.GET("/:id", func(ctx *gin.Context) {
			controller.GetRoleById(ctx)
		})
		roleGroup.DELETE("/:id", func(ctx *gin.Context) {
			controller.DeleteRole(ctx)
		})
	}
}
