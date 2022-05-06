package handlers

import (
	"net/http"

	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/team/vpn"
	"github.com/gin-gonic/gin"
)

func GetVpnConfigHandler(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}
	username := metadata.UserId
	congig, err := vpn.CreateVpnCLient(username).DownloadVpnConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.Data(200, "plain/text; charset=utf-8", []byte(congig))

}
