package mongo

import (
	"context"
	"errors"
	"tugas4go/database"
	modelMongo "tugas4go/app/model/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCol *mongo.Collection

func InitUserCollection() {
	userCol = database.MongoDB.Collection("users")
}

func getUserCol() (*mongo.Collection, error) {
	if userCol == nil {
		return nil, errors.New("user collection belum diinisialisasi, panggil InitUserCollection() dulu")
	}
	return userCol, nil
}

func CreateUser(user modelMongo.UserMongo) error {
	col, err := getUserCol()
	if err != nil {
		return err
	}
	_, err = col.InsertOne(context.TODO(), user)
	return err
}

func FindUserByEmail(email string) (*modelMongo.UserMongo, error) {
	col, err := getUserCol()
	if err != nil {
		return nil, err
	}
	var user modelMongo.UserMongo
	err = col.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByToken(token string) (*modelMongo.UserMongo, error) {
	col, err := getUserCol()
	if err != nil {
		return nil, err
	}
	var user modelMongo.UserMongo
	err = col.FindOne(context.TODO(), bson.M{"token": token}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserToken(ctx context.Context, userID primitive.ObjectID, token string) error {
	col, err := getUserCol()
	if err != nil {
		return err
	}

	_, err = col.UpdateOne(ctx,
		bson.M{"_id": userID},
		bson.M{"$set": bson.M{"token": token}},
	)
	return err
}