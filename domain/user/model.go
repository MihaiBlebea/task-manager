package user

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
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(username, email, password string) (*User, error) {
	u := &User{Username: username, Email: email}

	hash, err := u.hashPassword(password)
	if err != nil {
		return u, err
	}

	u.Password = hash

	return u, nil
}

func (u *User) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}

func (u *User) GenerateJWT() (string, error) {
	if u.ID == 0 {
		return "", errors.New("No valid userID")
	}

	secretKey := os.Getenv("AUTH_SECRET")
	if secretKey == "" {
		return "", errors.New("Please provide a secret key")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  u.ID,
		"username": u.Username,
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tkn, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	u.Token = tkn

	return tkn, nil
}
