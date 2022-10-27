package models

import (
	"time"
)

type Entity struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Name      string    `bson:"name"`
	Login     string    `bson:"login"`
	Hash      string    `bson:"hash"`
	Subnet    string    `bson:"subnet"`
	IP        string    `bson:"ip"`
	Visible   bool      `bson:"visible"`
	Blocked   bool      `bson:"blocked"`
}
type EntityInfo struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Name      string    `bson:"name" json:"name"`
	Login     string    `bson:"login" json:"login"`
	Subnet    string    `bson:"subnet" json:"subnet"`
	IP        string    `bson:"ip" json:"ip"`
	Visible   bool      `bson:"visible" json:"visible"`
	Blocked   bool      `bson:"blocked" json:"blocked"`
}
