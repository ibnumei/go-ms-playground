package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ibnumei/go-ms-playground/internal/app/domain"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{userRepo}
}

var signatureKey = []byte("mySignaturePrivateKey")

func (us UserService) Register(ctx context.Context, userBody domain.User) (string, error) {
	// userObject := userBody
	// fmt.Println("userObject UserService", userObject)
	// user := GenerateNewUser(&userObject)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	userBody.Password = string(hashedPassword)
	fmt.Println("UserService", userBody)

	if err := us.userRepository.Create(ctx, &userBody); err != nil {
		return "", errors.New("Failed Create User")
	}

	return generateJWT(userBody.ID, userBody.Username)
}

func generateJWT(userID int, Username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iss":      "edspert",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signatureKey)
	if err != nil {
		return "", err
	}
	return stringToken, nil
}
