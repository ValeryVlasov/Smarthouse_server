package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/ValeryVlasov/Smarthouse_server"
	"github.com/ValeryVlasov/Smarthouse_server/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"time"
)

const (
	salt       = "adivh4hec1328hchiucshadi"
	signingKey = "j9239M(Z)Eumwfdychsa"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user Smarthouse_server.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) IsSameUser(login, password interface{}) (Smarthouse_server.User2, bool) {
	user, err := s.repo.GetUser2(cast.ToString(login), generatePasswordHash(cast.ToString(password)))
	if err != nil {
		fmt.Println(err.Error())
		return user, false
	}
	if user.Username != cast.ToString(login) || user.Password_hash != cast.ToString(generatePasswordHash(cast.ToString(password))) {
		return user, false
	}
	return user, true
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
