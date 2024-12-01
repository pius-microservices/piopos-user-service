package interfaces

import (
	"github.com/pius-microservices/piopos-user-service/package/database/models"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	SignUp(userData *models.User) (*models.User, error)
	UpdateUser(userData *models.User) (*models.User, error)

	GetUsers() (*models.Users, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserService interface {
	SignUp(data *models.User) (gin.H, int)
	VerifyAccount(email string, otp string) (gin.H, int)
	SendNewOTPCode(email string) (gin.H, int)
	// UpdateUser(userData *models.User) (*models.User, error)

	GetUsers() (gin.H, int)
	GetUserById(id string) (gin.H, int)
	GetUserByEmail(email string) (gin.H, int)
}
