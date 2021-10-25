package component

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}
