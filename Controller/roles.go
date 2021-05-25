package controller

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/sen329/test5/Model"

	"github.com/dgrijalva/jwt-go"
)

var user model.User

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return JwtKey, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
