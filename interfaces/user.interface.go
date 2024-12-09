package interfaces

import (
	"github.com/pius-microservices/piopos-user-service/package/database/models"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	SignUp(userData *models.User) (*models.User, error)
	UpdateUser(userData *models.User) (*models.User, error)
	UpdateUserProfile(userData *models.User, id string) (*models.User, error)
	UpdatePassword(id string, password string) (*models.User, error)
	CreateRefreshToken(refreshToken *models.RefreshToken) (*models.RefreshToken, error)
	ValidateRefreshToken(userId string, refreshToken string) (*models.RefreshToken, error)

	GetUsers() (*models.Users, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserService interface {
	SignUp(data *models.User) (gin.H, int)
	VerifyAccount(email string, otp string) (gin.H, int)
	SendNewOTPCode(email string) (gin.H, int)
	UpdateUserProfile(userData *models.User, id string) (gin.H, int)
	UpdatePassword(id string, password string) (gin.H, int)
	CreateRefreshToken(userId string) (gin.H, int)
	ValidateRefreshToken(userId string, refreshToken string) (gin.H, int)

	GetUsers() (gin.H, int)
	GetUserById(id string) (gin.H, int)
	GetUserByEmail(email string) (gin.H, int)
}
