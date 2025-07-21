package mongodb

import (
	"context"
	"os"

	"github.com/Alan-Gomes1/go-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MONGODB_URL  = "MONGODB_URL"
	MONGODB_NAME = "MONGODB_NAME"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoDBUri := os.Getenv(MONGODB_URL)
	mongoDBName := os.Getenv(MONGODB_NAME)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBUri))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	logger.Info("Connected to MongoDB!")
	dataBase := client.Database(mongoDBName)
	return dataBase, nil
}
