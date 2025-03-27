package helpers

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once    sync.Once
	db      *mongo.Database
	initErr error
)

func initMongo() {
	_ = godotenv.Load("../../.env.test")

	uri := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		initErr = logError("failed to connect to mongo: %v", err)
		return
	}

	db = client.Database(dbName)
}

func CleanupCollection(collectionName string) {
	once.Do(initMongo)

	if initErr != nil || db == nil {
		log.Printf("[Cleanup] Skipping cleanup, Mongo not initialized: %v", initErr)
		return
	}

	log.Printf("[Cleanup] Cleaning collection: %s", collectionName)

	_, err := db.Collection(collectionName).DeleteMany(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Printf("[Cleanup] Failed to cleanup collection %s: %v", collectionName, err)
	}
}

func logError(format string, args ...interface{}) error {
	err := log.Output(2, "[Cleanup] "+format)
	if len(args) > 0 {
		log.Printf("[Cleanup] "+format, args...)
	}
	return err
}
