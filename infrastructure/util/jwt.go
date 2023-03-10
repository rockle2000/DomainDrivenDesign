package util

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"DDD_Project/infrastructure/config"
)

var (
	TokenExpired     error = errors.New("token is expired")
	TokenNotValidYet error = errors.New("token not active yet")
	TokenMalformed   error = errors.New("token malformed")
	TokenInvalid     error = errors.New("invalid token")
)

type JWT struct {
	SigningKey []byte
}
type CustomClaims struct {
	CustomerId int    `json:"customer_id"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	jwt.RegisteredClaims
}

type TokenInfo struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func NewJWT(config *config.AppConfig) JWT {
	return JWT{
		SigningKey: []byte(config.Secret),
	}
}

func (j *JWT) GenerateToken(config *config.AppConfig, customerId int, email, name string) (TokenInfo, error) {
	cfExpireTime := config.TokenExpired
	currentTime := time.Now()
	expireTime := time.Duration(cfExpireTime)
	tokenExpireTime := currentTime.Add(expireTime * time.Second)

	claims := CustomClaims{
		CustomerId: customerId,
		Email:      email,
		Name:       name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpireTime),
		},
	}
	token, err := j.CreateToken(claims)
	return TokenInfo{
		AccessToken:  token,
		RefreshToken: "",
		ExpiredAt:    tokenExpireTime,
	}, err
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}

}

func (j *JWT) ExtractToken(header string) (string, error) {
	tokenHeader := strings.Split(header, " ")
	if tokenHeader[0] != "Bearer" || len(tokenHeader) < 2 || strings.TrimSpace(tokenHeader[1]) == "" {
		return "", errors.New("invalid token")
	}
	return tokenHeader[1], nil
}
