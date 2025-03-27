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

	if notes == nil {
		notes = []Note{}
	}

	return notes, nil
}

func (r *MongoNoteRepository) UpdateByID(noteID string, userID string, updates map[string]interface{}) (Note, error) {
	noteObjectID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		return Note{}, errors.New("invalid note id")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return Note{}, errors.New("invalid user id")
	}

	updates["updated_at"] = time.Now()

	filter := bson.M{"_id": noteObjectID, "user_id": userObjectID}
	update := bson.M{"$set": updates}

	res := r.collection.FindOneAndUpdate(
		context.Background(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	)

	var updatedNote Note
	err = res.Decode(&updatedNote)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Note{}, errors.New("note not found or not owned by user")
		}
		return Note{}, err
	}

	return updatedNote, nil
}

func (r *MongoNoteRepository) DeleteByID(noteID string, userID string) error {
	noteObjectID, err := primitive.ObjectIDFromHex(noteID)
	if err != nil {
		return errors.New("invalid note id")
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id")
	}

	res, err := r.collection.DeleteOne(context.Background(), bson.M{
		"_id":     noteObjectID,
		"user_id": userObjectID,
	})
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("note not found or not owned by user")
	}

	return nil
}
