package user

import (
	"time"

	"github.com/pius-microservices/piopos-user-service/interfaces"
	"github.com/pius-microservices/piopos-user-service/package/database/models"
	"github.com/pius-microservices/piopos-user-service/package/utils"

	"github.com/gin-gonic/gin"
)

type userService struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *userService {
	return &userService{repo}
}

func (service *userService) SignUp(userData *models.User) (gin.H, int) {
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	otpCode := utils.GenerateOTP(6)

	userData.Username = utils.GenerateUsername(userData.Email)
	userData.Password = hashedPassword
	userData.OTPCode = otpCode
	userData.OTPExpiration = time.Now().Add(30 * time.Minute)

	newData, err := service.repo.SignUp(userData)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			return gin.H{"status": 409, "message": "Email is already used"}, 409
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"data": newData}, 201
}

func (service *userService) VerifyAccount(email string, otp string) (gin.H, int) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}

		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	if user.OTPCode != otp {
		return gin.H{"status": 401, "message": "Invalid OTP code"}, 401
	}

	if time.Now().After(user.OTPExpiration) {
		return gin.H{"status": 401, "message": "OTP code has expired"}, 401
	}

	user.IsVerified = true
	user.OTPCode = ""
	user.OTPExpiration = time.Time{}

	_, err = service.repo.UpdateUser(user)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"message": "Account verified successfully"}, 200
}

func (service *userService) SendNewOTPCode(email string) (gin.H, int) {
	user, err := service.repo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}

		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	otpCode := utils.GenerateOTP(6)
	user.OTPCode = otpCode
	user.OTPExpiration = time.Now().Add(30 * time.Minute)

	_, err = service.repo.UpdateUser(user)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"message": "New OTP code sent successfully"}, 200
}

func (service *userService) UpdateUserProfile(userData *models.User, id string) (gin.H, int) {
	_, err := service.repo.GetUserById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}

		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	updatedUser, err := service.repo.UpdateUserProfile(userData, id)

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "User updated successfully", "data": &updatedUser}, 200
}

func (service *userService) UpdatePassword(id string, password string) (gin.H, int) {
	_, err := service.repo.GetUserById(id)

	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	_, err = service.repo.UpdatePassword(id, hashedPassword)

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Password updated successfully"}, 200
}

func (service *userService) GetUsers() (gin.H, int) {
	users, err := service.repo.GetUsers()

	if err != nil {
		return gin.H{"status": 404, "message": err.Error()}, 404
	}

	return gin.H{"status": 200, "message": "All users fetched successfully", "data": users}, 200
}

func (service *userService) GetUserById(id string) (gin.H, int) {
	user, err := service.repo.GetUserById(id)

	if err != nil {
		return gin.H{"status": 500, "message": "Failed to retrieve user data"}, 500
	}

	return gin.H{"status": 200, "data": user}, 200
}

func (service *userService) GetUserByEmail(email string) (gin.H, int) {
	user, err := service.repo.GetUserByEmail(email)

	if err != nil {
		if err.Error() == "record not found" {
			return gin.H{"status": 404, "message": "User not found"}, 404
		}

		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "data": user}, 200
}
