package database

import (
	"context"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/team/models"
	"go.mongodb.org/mongo-driver/bson"
)

func AddTeam(t *models.Team) error {
	col := database.GetCollection(database.DB, "teams")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return err
	}

	return nil
}

func GetTeam(login string) (*models.TeamInfo, error) {
	col := database.GetCollection(database.DB, "teams")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var team models.TeamInfo
	filter := bson.M{"login": login}
	err := col.FindOne(ctx, filter).Decode(&team)
	if err != nil {
		return nil, err
	}
	return &team, err

}
