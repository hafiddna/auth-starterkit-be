package repository

import (
	"context"
	"github.com/hafiddna/auth-starterkit-be/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserProfileRepository interface {
	FindOneByUserID(id string) interface{}
}

type userProfileRepository struct {
	db *mongo.Database
}

func (r *userProfileRepository) environment() *mongo.Collection {
	return r.db.Collection("user_profiles")
}

func NewUserProfileRepository(db *mongo.Database) UserProfileRepository {
	return &userProfileRepository{db: db}
}

func (r *userProfileRepository) FindOneByUserID(id string) interface{} {
	var userProfile entity.UserProfile

	err := r.environment().FindOne(context.TODO(), bson.M{"user_id": id}).Decode(&userProfile)
	if err != nil {
		return nil
	}

	return userProfile
}