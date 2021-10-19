package src

import (
  "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UserInfo struct {
  ID primitive.ObjectID `bson:"_id,omitempty"`
  Username string `bson:"username,omitempty"`
  Pin      int    `bson:"pin,omitempty"`
  Amount   float64 `bson:"amount,omitempty"`
  Profession string `bson:""`
}

type Transaction struct {
  ID primitive.ObjectID `bson:"_id,omitempty"`
  Type string `bson:"transactionType,omitempty"`
  Amount float64 `bson:"amount,omitempty"`
  User string `bson:"username,omitempty"`
  Date float64 `bson:"timestamp,omitempty"`
}
