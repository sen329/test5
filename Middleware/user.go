package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	_ "github.com/go-sql-driver/mysql"

	controller "github.com/sen329/test5/Controller"
	model "github.com/sen329/test5/Model"
)

var db *sql.DB

var check bool

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]

			token, err := jwt.ParseWithClaims(jwtToken, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(controller.JwtKey), nil
			})

			if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
				ctx := context.WithValue(context.Background(), "user_id", claims.User_id)
				ctx = context.WithValue(ctx, "role_id", claims.Role_id)

				next.ServeHTTP(w, r.WithContext(ctx))

			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}
