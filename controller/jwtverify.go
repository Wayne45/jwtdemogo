package controller

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

var publicKey *rsa.PublicKey

const bearerPrefix = "Bearer "

func init() {
	bytes, err := os.ReadFile("./files/publickey.pem")
	if err != nil {
		panic(err)
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
}

// Verify JWT token and print the claims
// @Summary Verify JWT token
// @Schemes
// @Description Verify JWT token
// @Tags JWT
// @Produce json
// @Success 200 {string} JWT token
// @Router /jwt-tokens [get]
func VerifyJwt(g *gin.Context) {
	authorization := g.GetHeader("Authorization")
	if strings.HasPrefix(authorization, bearerPrefix) {
		tokenString := authorization[len(bearerPrefix):]
		token, err := jwt.Parse(tokenString, func(tok *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})
		if err != nil {
			g.JSON(http.StatusUnauthorized, HttpError{ErrBadAuth, err.Error()})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				g.JSON(http.StatusOK, claims)
			}
		}
	} else {
		g.JSON(http.StatusUnauthorized, HttpError{ErrMissingAuth, ""})
	}
}
