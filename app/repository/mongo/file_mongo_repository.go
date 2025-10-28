package mongo

import (
	"context"
	"time"

	modelMongo "tugas4go/app/model/mongo"
	"tugas4go/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var fileCol *mongo.Collection

func InitFileCollection() {
	fileCol = database.MongoDB.Collection("uploaded_files")
}

func CreateFileMongo(file modelMongo.FileMongo) error {
	file.UploadedAt = time.Now()
	_, err := fileCol.InsertOne(context.TODO(), file)
	return err
}

func GetAllFilesMongo() ([]modelMongo.FileMongo, error) {
	var results []modelMongo.FileMongo
	cur, err := fileCol.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var f modelMongo.FileMongo
		cur.Decode(&f)
		results = append(results, f)
	}
	return results, nil
}

// GET BY FILENAME (tambahan opsional untuk ambil metadata file tertentu)
func GetFileByName(filename string) (*modelMongo.FileMongo, error) {
	var file modelMongo.FileMongo
	err := fileCol.FindOne(context.TODO(), bson.M{"file_name": filename}).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

// DELETE FILE RECORD 
func DeleteFileRecord(filename string) error {
	_, err := fileCol.DeleteOne(context.TODO(), bson.M{"file_name": filename})
	return err
}
