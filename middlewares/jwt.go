package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
	UserRole string `json:"userRole"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, userName string, userRole string) string {
	var userClaims = jwtCustomClaims{
		userId, userName, userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	resultJWT, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return resultJWT
}
