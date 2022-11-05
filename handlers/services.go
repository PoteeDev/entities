package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/admin/models"
	"github.com/PoteeDev/auth/auth"
	scoreModels "github.com/PoteeDev/scores/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Reputation  int     `json:"reputation"`
	Lost        int     `json:"lost"`
	Gained      int     `json:"gained"`
	Status      int     `json:"status"`
	SLA         float64 `json:"sla"`
}

func GetServices(c *gin.Context) {
	var scenario models.Scenario
	var scoreboard scoreModels.Score
	var services []Service
	metadata, err := auth.NewToken().ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"detail": "unauthorized"})
		return
	}
	col := database.GetCollection(database.DB, "settings")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = col.FindOne(ctx, bson.M{"id": "scenario"}).Decode(&scenario)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err})
		return
	}
	scoreCol := database.GetCollection(database.DB, "scoreboard")
	emtyScore := false
	err = scoreCol.FindOne(ctx, bson.M{"id": metadata.UserId}).Decode(&scoreboard)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			emtyScore = true
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"detail": err})
			return
		}
	}
	for name, serviceInfo := range scenario.Services {
		service := Service{
			Name:        serviceInfo.Name,
			Description: serviceInfo.Description,
			Reputation:  serviceInfo.Reputation,
			Gained:      0,
			Lost:        0,
			SLA:         -1,
			Status:      -1,
		}
		log.Println(scoreboard)
		if !emtyScore {
			service.Reputation = scoreboard.Srv[name].Reputation
			service.Lost = scoreboard.Srv[name].Lost
			service.Gained = scoreboard.Srv[name].Gained
			service.Status = scoreboard.Srv[name].Status
			service.SLA = scoreboard.Srv[name].SLA
		}
		services = append(services, service)
	}

	c.JSON(http.StatusOK, gin.H{"services": services})
}
