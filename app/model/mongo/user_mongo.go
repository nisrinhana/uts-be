package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserMongo struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Role     string             `bson:"role" json:"role"`
	Token    string             `bson:"token" json:"token"`
}
