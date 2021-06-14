package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Email   string `json:"email"`
	User_id string `json:"user_id"`
	Role_id string `json:"role_id"`
	jwt.StandardClaims
}
