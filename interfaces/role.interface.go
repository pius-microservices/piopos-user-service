package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/pius-microservices/piopos-user-service/package/database/models"
)

type RoleRepo interface {
	AddRole(data *models.Role) (*models.Role, error)
	GetRoles() (*models.Roles, error)
	GetRoleById(id string) (*models.Role, error)
	DeleteRole(id string) error
}

type RoleService interface {
	AddRole(data *models.Role) (gin.H, int)
	GetRoles() (gin.H, int)
	GetRoleById(id string) (gin.H, int)
	DeleteRole(id string) (gin.H, int)
}