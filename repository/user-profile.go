package repository

import (
	"context"
	"github.com/hafiddna/auth-starterkit-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserProfileRepository interface {
	FindOneByUserID(id string) (data model.UserProfile, err error)
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

func (r *userProfileRepository) FindOneByUserID(id string) (data model.UserProfile, err error) {
	var userProfile model.UserProfile

	err = r.environment().FindOne(context.TODO(), bson.M{"user_id": id}).Decode(&userProfile)
	if err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
