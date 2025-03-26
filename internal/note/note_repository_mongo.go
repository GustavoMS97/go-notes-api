package note

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoNoteRepository struct {
	collection *mongo.Collection
}

func NewMongoNoteRepository(db *mongo.Database) *MongoNoteRepository {
	return &MongoNoteRepository{
		collection: db.Collection("notes"),
	}
}

func (r *MongoNoteRepository) Create(note Note) (Note, error) {
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(context.Background(), note)
	if err != nil {
		return Note{}, err
	}

	note.ObjectID = res.InsertedID.(primitive.ObjectID)

	return note, nil
}

func (r *MongoNoteRepository) FindAllByUserID(userID string, search string) ([]Note, error) {
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user id")
	}

	filter := bson.M{"user_id": objectID}

	if search != "" {
		filter["title"] = bson.M{
			"$regex":   search,
			"$options": "i",
		}
	}

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var notes []Note
	for cursor.Next(context.Background()) {
		var note Note
		if err := cursor.Decode(&note); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
