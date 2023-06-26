package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Controllers struct {
	MongoDB *mongo.Database
}
