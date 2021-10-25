package component

import (
	"github.com/alicfeng/gin_template/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// MakeToken */
func MakeToken(uid int) (string, error) {
	// 24小时候过期
	expireTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		uid,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "rx_service",
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Jwt.Secret))

	return token, err
}

// ParseToken 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	//
	return nil, err
}

// GetUserId 获取用户标识 /**
func GetUserId(context *gin.Context) int {

	return context.GetInt("user_id")
}
