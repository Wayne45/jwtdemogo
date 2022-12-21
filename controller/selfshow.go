package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

// Show identity by extracting from header
// @Summary Show identity
// @Schemes
// @Description Show identity
// @Tags JWT
// @Produce json
// @Success 200 {string} Identity string
// @Router /self [get]
func ShowSelf(g *gin.Context) {
	dump, _ := httputil.DumpRequest(g.Request, true)
	fmt.Println(string(dump))

	self := g.GetHeader("x-jwt-claim-nested-claim")
	g.JSON(http.StatusOK, gin.H{
		"identity": self,
	})
}
