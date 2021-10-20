package src

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
)

var (
	amountTooBigErr = errors.New("Amount is larger than amount in bank")
	invalidOption   = errors.New("Invalid option")
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

func (cm *CashMachine) CreateAccount(ctx context.Context) bool {

	fmt.Println("Choose a username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Choose a pin: ")
	var pin int
	fmt.Scanln(&pin)

	fmt.Println("Choose a profession (developer, engineer, doctor, scientist): ")
	var profession string
	fmt.Scanln(&profession)

	var amount float64
	switch profession {
	case "developer":
		amount = 50000
	case "engineer":
		amount = 30000
	case "doctor":
		amount = 70000
	case "scientist":
		amount = 75000
	}

	userInfo := UserInfo{
		Username:   username,
		Pin:        pin,
		Amount:     amount,
		Profession: profession,
	}

	_, err := cm.client.Collection("users").InsertOne(ctx, userInfo)

	if err != nil {
		return false
	}

	return true
}

func (cm *CashMachine) Login(ctx context.Context) bool {

	fmt.Println("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Enter a pin: ")
	var pin int
	fmt.Scanln(&pin)

	var userInfo *UserInfo

	res, err := cm.client.Collection("users").Find(ctx, bson.M{"username": username, "pin": pin})

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

func (cm *CashMachine) WithdrawAdd(ctx context.Context, txType string) error {

	fmt.Println("Enter amount: ")
	var amount float64
	_, err := fmt.Scanln(&amount)
	if err != nil {
		log.Fatalf("ATM - Withdaw - Error %v", err)
		return err
	}

	if amount > cm.info.Amount {
		return amountTooBigErr
	}

	switch txType {
	case "withdraw":
		cm.info.Amount -= amount
		break
	case "add":
		cm.info.Amount += amount
		break
	}

	_, err = cm.client.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": cm.info.ID},
		bson.D{
			{"$set", bson.D{
				{"amount", cm.info.Amount},
			}},
		})

	if err != nil {
		log.Fatalf("ATM - Withdraw - Database error %v", err)
		return err
	}

	return nil
}
func (cm *CashMachine) Gamble(ctx context.Context) error {

	fmt.Println(`


	Odds: 51 win / 49 lose , 2x multiplier
	- 1 - Start game


`)

	fmt.Println("> ")
	var option int
	fmt.Scanln(&option)
	if option != 1 {
		return invalidOption
	}

	min := 0
	max := 100

	result := rand.Intn(max-min) + min

	if result >= 51 {
		cm.info.Amount *= 2
	} else {
		cm.info.Amount /= 2
	}

	_, err := cm.client.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": cm.info.ID},
		bson.D{
			{"$set", bson.D{
				{"amount", cm.info.Amount},
			}},
		})

	if err != nil {
		log.Fatalf("ATM - Gambling - Database error %v", err)
		return err
	}

	return nil

}
func (cm *CashMachine) Work(ctx context.Context) bool {
	var hourlyWage float64
	var workingHours float64

	switch cm.info.Profession {
	case "developer":
		hourlyWage = 35.50
		workingHours = 8
	case "engineer":
		hourlyWage = 30.50
		workingHours = 7
	case "doctor":
		hourlyWage = 40
		workingHours = 10
	case "scientist":
		hourlyWage = 45
		workingHours = 10
	}

	newAmount := hourlyWage * workingHours
	cm.info.Amount += newAmount

	_, err := cm.client.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": cm.info.ID},
		bson.D{
			{"$set", bson.D{
				{"amount", cm.info.Amount},
			}},
		})

	if err != nil {
		log.Fatalf("ATM - Work - Database error %v", err)
		return false
	}

	return true

}
