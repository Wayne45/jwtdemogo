package controller

import (
	"crypto/rsa"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

var privateKey *rsa.PrivateKey

func init() {
	bytes, err := os.ReadFile("/files/privatekey.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		panic(err)
	}
}

// SignJwt Sign JWT token for testing
// @Summary Sign JWT token
// @Schemes
// @Description Sign JWT token
// @Tags JWT
// @Produce json
// @Success 200 {string} JWT token
// @Router /jwt-tokens [post]
func SignJwt(g *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		// now + 1h
		"exp": time.Now().Add(time.Hour).UnixMilli() / 1000,
		"iss": "jwtdemogo",
	})
	signed, err := token.SignedString(privateKey)
	if err != nil {
		g.JSON(http.StatusInternalServerError, HttpError{ErrJwtSignError, err.Error()})
	} else {
		g.JSON(http.StatusOK, gin.H{"token": signed})
	}
}
