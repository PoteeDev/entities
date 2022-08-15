package handlers

import (
	"fmt"
	"net/http"

	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/entities/vpn"
	"github.com/gin-gonic/gin"
)

func GenerateVpnConfig(c *gin.Context) {
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}
	err = vpn.CreateVpnCLient(metadata.UserId).CreateConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("config for %s created", metadata.UserId)})
}

func GetVpnConfig(c *gin.Context) {
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
