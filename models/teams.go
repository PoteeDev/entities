package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Name      string             `bson:"name"`
	Login     string             `bson:"login"`
	Hash      string             `bson:"hash"`
	Visible   bool               `bson:"visible"`
	Blocked   bool               `bson:"blocked"`
}
type TeamInfo struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Name      string    `bson:"name" json:"name"`
	Login     string    `bson:"login" json:"login"`
	Visible   bool      `bson:"visible" json:"visible"`
	Blocked   bool      `bson:"blocked" json:"blocked"`
}
