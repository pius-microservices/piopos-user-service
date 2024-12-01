package role

import (
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/pius-microservices/piopos-user-service/interfaces"
	"github.com/pius-microservices/piopos-user-service/package/database/models"
)

type roleController struct {
	service interfaces.RoleService
}

func NewController(service interfaces.RoleService) *roleController {
	return &roleController{service}
}

// AddRole godoc
// @Summary Add Role
// @Description Create new role
// @Tags Roles
// @Accept json
// @Produce json
// @Param userData body models.AddRoleRequest true "User data"
// @Success 201
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /api/role [post]
func (controller *roleController) AddRole(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	var roleData models.Role

	err := ctx.ShouldBindJSON(&roleData)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	_, err = govalidator.ValidateStruct(&roleData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	responseData, status := controller.service.AddRole(&roleData)

	ctx.JSON(status, responseData)
}

// GetRoles godoc
// @Summary Get roles
// @Description Get all user roles
// @Tags Roles
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/role [get]
func (controller *roleController) GetRoles(ctx *gin.Context) {
	responseData, status := controller.service.GetRoles()
	ctx.JSON(status, responseData)
}

// GetRoleById godoc
// @Summary Get a role
// @Description Get role by role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /api/user-service/role/{id} [get]
func (controller *roleController) GetRoleById(ctx *gin.Context) {
	id := ctx.Param("id")

	responseData, status := controller.service.GetRoleById(id)

	ctx.JSON(status, responseData)
}

// DeleteRole godoc
// @Summary Delete a role
// @Description Delete a role by role ID
// @Tags Roles
// @Accept json
// @Produce json
// @Param id path string true "Role ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/user-service/role/{id} [delete]
func (controller *roleController) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")

	responseData, status := controller.service.DeleteRole(id)

	ctx.JSON(status, responseData)
}