package note

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
