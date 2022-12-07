package domain

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	NoHp      string    `json:"noHp"`
	IsActive  bool      `json:"isActive"`
	CreatedBy string    `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedBy string    `json:"updatedBy"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// semula ingin di tempatkan di user_service, karena masih berbentuk function
// tapi tidak bisa karena menggunakan method struct user "func (user User)"
func (user User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

var privateKey = os.Getenv("SIGNATURE_KEY")
var signatureKey = []byte(privateKey)

//decrypt jwt untuk auth akses data, middleware request
func (u User) DecryptJWT(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		} 
		return signatureKey, nil
	})

	data := make(map[string]interface{})
	if err != nil {
		return data, err
	}
	if !parsedToken.Valid {
		return data, errors.New("invalid token")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
