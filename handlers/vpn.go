package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/auth/auth"
	"github.com/PoteeDev/entities/models"
	"github.com/PoteeDev/entities/vpn"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GenerateVpnConfig(c *gin.Context) {
	var entity models.Entity
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}

	// TODO get ip of entity
	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = col.FindOne(ctx, bson.M{"login": metadata.UserId}).Decode(&entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}
	if entity.Subnet == "" {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "can not generate config"})
		return
	}

	vpnClient := vpn.CreateVpnCLient(metadata.UserId)
	err = vpnClient.CreateConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}

	err = vpnClient.AddRoute(entity.Subnet)
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
