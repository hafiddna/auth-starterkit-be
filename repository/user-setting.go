package repository

import (
	"context"
	"github.com/hafiddna/auth-starterkit-be/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSettingRepository interface {
	FindOneByUserID(id string) (data model.UserSetting, err error)
}

type userSettingRepository struct {
	db *mongo.Database
}

func (r *userSettingRepository) environment() *mongo.Collection {
	return r.db.Collection("user_settings")
}

func NewUserSettingRepository(db *mongo.Database) UserSettingRepository {
	return &userSettingRepository{db: db}
}

func (r *userSettingRepository) FindOneByUserID(id string) (data model.UserSetting, err error) {
	var userSetting model.UserSetting

	err = r.environment().FindOne(context.TODO(), bson.M{"user_id": id}).Decode(&userSetting)
	if err != nil {
		return userSetting, err
	}

	return userSetting, nil
}
