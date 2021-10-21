package src

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInfo struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `bson:"username,omitempty"`
	Pin        int                `bson:"pin,omitempty"`
	Amount     float64            `bson:"amount,omitempty"`
	Profession string             `bson:"profession,omitempty"`
}
