package user

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{
		collection: db.Collection("users"),
	}
}

func (r *MongoUserRepository) Create(user User) (User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	res, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		return User{}, err
	}

	user.ObjectID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (r *MongoUserRepository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) FindByID(id string) (*User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid user id format")
	}

	var user User
	err = r.collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
