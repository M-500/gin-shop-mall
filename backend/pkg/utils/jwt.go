package utils

import (
	"backend/internal/config"
	e "backend/pkg/constant"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024/1/9 11:53
type JWT struct {
	SigningKey []byte
}

// CustomClaims
// @Description: Jwt相关数据
type CustomClaims struct {
	ID          uint
	Email       string
	UserName    string
	AuthorityId uint
	jwt.StandardClaims
}

func NewJwt() *JWT {
	jwtCfg := config.ConfigInstance.Jwt
	return &JWT{
		[]byte(jwtCfg.SecretKey), // 可以设置过期时间
	}
}
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(e.GetMessage(e.TokenMalformed))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New(e.GetMessage(e.TokenExpired))
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(e.GetMessage(e.TokenNotValidYet))
			} else {
				return nil, errors.New(e.GetMessage(e.TokenInvalid))
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, errors.New(e.GetMessage(e.TokenInvalid))
}

func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", errors.New(e.GetMessage(e.TokenInvalid))
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
