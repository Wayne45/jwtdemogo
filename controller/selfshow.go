package controller

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Identity struct {
	Iss string `json:"iss"`
	Exp int32  `json:"exp"`
}

// Show identity by extracting from header
// @Summary Show identity
// @Schemes
// @Description Show identity
// @Tags JWT
// @Produce json
// @Success 200 {string} Identity string
// @Router /self [get]
func ShowSelf(g *gin.Context) {
	//dump, _ := httputil.DumpRequest(g.Request, true)
	//fmt.Println(string(dump))
	header := g.GetHeader("auth_user")
	if header != "" {
		decoded, err := base64.RawStdEncoding.DecodeString(header)
		if err != nil {
			g.JSON(http.StatusUnauthorized, HttpError{ErrBadAuth, err.Error()})
		} else {
			self := Identity{}
			err = json.Unmarshal(decoded, &self)
			if err != nil {
				g.JSON(http.StatusInternalServerError, HttpError{ErrServerError, err.Error()})
			} else {
				g.JSON(http.StatusOK, self)
			}
		}
	}
}
