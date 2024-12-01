package role

import (
	"github.com/gin-gonic/gin"
	"github.com/pius-microservices/piopos-user-service/interfaces"
	"github.com/pius-microservices/piopos-user-service/package/database/models"
)

type roleService struct {
	repo interfaces.RoleRepo
}

func NewService(repo interfaces.RoleRepo) *roleService {
	return &roleService{repo}
}

func (service *roleService) AddRole(roleData *models.Role) (gin.H, int) {
	newRole, err := service.repo.AddRole(roleData)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_roles_name\" (SQLSTATE 23505)" {
			return gin.H{"status": 409, "message": "Role name already exists"}, 409
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"data": newRole}, 201
}

func (service *roleService) GetRoles() (gin.H, int) {
	roles, err := service.repo.GetRoles()

	if err != nil {
		return gin.H{"status": 404, "message": err.Error()}, 404
	}

	return gin.H{"status": 200, "message": "All roles fetched successfully", "data": roles}, 200
}

func (service *roleService) GetRoleById(id string) (gin.H, int) {
	role, err := service.repo.GetRoleById(id)

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "data": role}, 200
}

func (service *roleService) DeleteRole(id string) (gin.H, int) {
	role := service.repo.DeleteRole(id)

	if role != nil {
		return gin.H{"status": 404, "message": "Role not found"}, 404
	}

	return gin.H{"status": 200, "message": "Role deleted successfully"}, 200
}
