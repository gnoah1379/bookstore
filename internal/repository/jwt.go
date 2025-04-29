package repository

import (
	"bookstore/internal/model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTRepo interface {
	GenerateJWT(ID int, username string) (string, error)
	VerifyJWT(tokenString string) (*model.UserClaims, error)
}

var _ JWTRepo = (*jwtRepo)(nil)

type jwtRepo struct {
	SecretKey []byte
}

func NewJWTRepo(secretKey string) JWTRepo {
	return &jwtRepo{SecretKey: []byte(secretKey)}
}

func (j *jwtRepo) GenerateJWT(ID int, username string) (string, error) {
	claims := model.UserClaims{
		Id:       ID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore",
			Subject:   "user-authentication",
			ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtRepo) VerifyJWT(tokenString string) (*model.UserClaims, error) {
	claims := &model.UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("phương thức ký không hợp lệ: %v\"", token.Header["alg"])
		}
		jwtSecret := j.SecretKey
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token không hợp lệ")
}
