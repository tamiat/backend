package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/tamiat/backend/pkg/domain/user"
)

func TokenVerifyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := fmt.Sprintf("%s", os.Getenv("JWT_SECRET"))
		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(secret), nil
			})
			if err != nil {
				//TODO
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				log.Println(claims)
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			//TODO
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
func GenerateToken(user user.User) (string, error) {
	var err error
	secret := fmt.Sprintf("%s", os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		//TODO change claims
		"email": user.Email,
		"iss":   "course",
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Signing method validation
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret signing key
		secret := fmt.Sprintf("%s", os.Getenv("JWT_SECRET"))
		return []byte(secret), nil
	})
}
