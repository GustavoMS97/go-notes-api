package helpers

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CleanupCollection(collectionName string) {
	log.Println("[Cleanup] Connecting to MongoDB for cleanup...", collectionName)
	if err := godotenv.Load("../../.env.test"); err != nil {
		log.Fatalf("failed to load .env.test: %v", err)
	}

	uri := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("failed to connect to mongo: %v", err)
	}

	collection := client.Database(dbName).Collection(collectionName)

	_, err = collection.DeleteMany(context.Background(), map[string]interface{}{})
	if err != nil {
		log.Fatalf("failed to cleanup collection %s: %v", collectionName, err)
	}
}
