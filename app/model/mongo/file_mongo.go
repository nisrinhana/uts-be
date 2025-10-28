package mongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileMongo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FileName   string             `bson:"file_name" json:"file_name"`
	FilePath   string             `bson:"file_path" json:"file_path"`
	FileType   string             `bson:"file_type" json:"file_type"`
	FileSize   int64              `bson:"file_size" json:"file_size"`
	AlumniID   string             `bson:"alumni_id" json:"alumni_id"` // <- ubah int jadi string
	UploadedBy string             `bson:"uploaded_by" json:"uploaded_by"`
	UploadedAt time.Time          `bson:"uploaded_at" json:"uploaded_at"`
}


type FileResponse struct {
	ID            string    `json:"id"`
	FileName      string    `json:"file_name"`
	OriginalName  string    `json:"original_name"`
	FilePath      string    `json:"file_path"`
	FileSize      int64     `json:"file_size"`
	FileType      string    `json:"file_type"`
	UploadedAt    time.Time `json:"uploaded_at"`
}