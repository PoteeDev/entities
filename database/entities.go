package database

import (
	"context"
	"time"

	"github.com/PoteeDev/admin/api/database"
	"github.com/PoteeDev/entities/models"
	"go.mongodb.org/mongo-driver/bson"
)

func AddEntity(t *models.Entity) error {
	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return err
	}

	return nil
}

func GetEntity(login string) (*models.EntityInfo, error) {
	col := database.GetCollection(database.DB, "entities")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var entity models.EntityInfo
	filter := bson.M{"login": login}
	err := col.FindOne(ctx, filter).Decode(&entity)
	if err != nil {
		return nil, err
	}
	return &entity, err
}
