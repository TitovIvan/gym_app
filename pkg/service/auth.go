package service

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zemags/gym_app/pkg/repository"
	gym_app "github.com/zemags/gym_app/store"
)

const tokenTTL = 12 * time.Hour

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user gym_app.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	user.CreatedAt = time.Now()
	// TODO: check if user is admin

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(name, password string) (string, error) {
	user, err := s.repo.GetUser(name, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(tokenTTL).Unix(),
				IssuedAt:  time.Now().Unix(),
			},
			user.ID,
		},
	)
	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))

}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("SALT"))))
}
