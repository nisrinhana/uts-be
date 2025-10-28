package mongo

import (
	"context"
	modelMongo "tugas4go/app/model/mongo"
	"tugas4go/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var pekerjaanCol *mongo.Collection

func InitPekerjaanCollection() {
	pekerjaanCol = database.MongoDB.Collection("pekerjaan_alumni")
}

// CREATE
func CreatePekerjaanMongo(p modelMongo.PekerjaanAlumniMongo) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	_, err := pekerjaanCol.InsertOne(context.TODO(), p)
	return err
}

// GET ALL
func GetAllPekerjaanMongo() ([]modelMongo.PekerjaanAlumniMongo, error) {
	var results []modelMongo.PekerjaanAlumniMongo
	cur, err := pekerjaanCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var p modelMongo.PekerjaanAlumniMongo
		cur.Decode(&p)
		results = append(results, p)
	}
	return results, nil
}

// GET BY ID
func GetPekerjaanByIDMongo(id string) (modelMongo.PekerjaanAlumniMongo, error) {
	var p modelMongo.PekerjaanAlumniMongo
	objID, _ := primitive.ObjectIDFromHex(id)
	err := pekerjaanCol.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&p)
	return p, err
}

// GET BY ALUMNI_ID
func GetPekerjaanByAlumniIDMongo(alumniID int) ([]modelMongo.PekerjaanAlumniMongo, error) {
	var results []modelMongo.PekerjaanAlumniMongo
	cur, err := pekerjaanCol.Find(context.TODO(), bson.M{"alumni_id": alumniID})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var p modelMongo.PekerjaanAlumniMongo
		cur.Decode(&p)
		results = append(results, p)
	}
	return results, nil
}

// UPDATE
func UpdatePekerjaanMongo(id string, update modelMongo.PekerjaanAlumniMongo) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	update.UpdatedAt = time.Now()

	_, err := pekerjaanCol.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.M{"$set": update},
	)
	return err
}

// DELETE
func DeletePekerjaanMongo(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := pekerjaanCol.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}
