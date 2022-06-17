package middleware

import (
	"BE_WEB_BEM_Proker/infrastructure/app"
	"BE_WEB_BEM_Proker/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateJWToken(id string) (string, error) {
	env, err := app.NewDriverApp()
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	signedToken, err := token.SignedString([]byte(env.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateJWToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		if bearerToken == "" {
			//respon
			return
		}
		bearerToken = strings.ReplaceAll(bearerToken, "Bearer ", "")
		token, err := jwt.Parse(bearerToken, ekstractToken)
		if err != nil {
			c.JSON(http.StatusForbidden, response.ResponseWhenFail("Failed to extract token", err.Error()))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("id", claims["id"])
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, response.ResponseWhenFail("Failed to extract token", err.Error()))
			return
		}
	}
}

func ekstractToken(token *jwt.Token) (interface{}, error) {
	env, err := app.NewDriverApp()
	if err != nil {
		return "", err
	}
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, jwt.ErrSignatureInvalid
	}
	return []byte(env.SecretKey), nil
}
