package src

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type CashMachine struct {
	info   *UserInfo
	client *mongo.Database
}

func (cm *CashMachine) NewConnection(ctx context.Context) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("ATM - Connection - Error %v", err)
	}

	cm.client = client.Database("ATM")
	log.Println("ATM - Connected!")
}
func (cm *CashMachine) Login(ctx context.Context) bool {

	var userInfo *UserInfo

	res, err := cm.client.Collection("users").Find(ctx, bson.M{"username": os.Getenv("USERNAME")})

	if err != nil {
		log.Fatalf("ATM - Login - Error %v", err)
		return false
	}

	err = res.Decode(&userInfo)
	if err != nil {
		log.Fatalf("ATM - Login - Error %v", err)
		return false
	}
	return true
}

func (cm *CashMachine) Withdraw() {}
func (cm *CashMachine) Add()      {}
func (cm *CashMachine) Gamble()   {}
func (cm *CashMachine) Work()     {}
