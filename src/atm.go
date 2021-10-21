1package src

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Every function has a method receiver as a pointer to CashMachine for easy use in the struct

// globally defined errors that I can use as they come up more than once
var (
	amountTooBigErr = errors.New("Amount is larger than amount in bank")
	invalidOption   = errors.New("Invalid option")
)

// Struct containing all the functions and info regarding the user and the cash machine
type CashMachine struct {
	info   *UserInfo
	client *mongo.Database
}

func (cm *CashMachine) NewConnection(ctx context.Context) {
	// New Connection creates a new connection to the database
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("ATM - Connection - Error %v", err)
	}

	cm.client = client.Database("ATM")
	log.Println("ATM - Connected!")
}

func (cm *CashMachine) CreateAccount(ctx context.Context) bool {

	// User inputs
	fmt.Println("Choose a username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Choose a pin: ")
	var pin int
	fmt.Scanln(&pin)

	fmt.Println("Choose a profession (developer, engineer, doctor, scientist): ")
	var profession string
	fmt.Scanln(&profession)

	// Starting amount of money
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

	// Creates new valid user in the database
	_, err := cm.client.Collection("users").InsertOne(ctx, userInfo)

	if err != nil {
		return false
	}

	return true
}

func (cm *CashMachine) Login(ctx context.Context) bool {

	// User inputs
	fmt.Println("Enter username: ")
	var username string
	fmt.Scanln(&username)

	fmt.Println("Enter a pin: ")
	var pin int
	fmt.Scanln(&pin)

	// Checks if the user is in the database, if not then it returns an error
	if err := cm.client.Collection("users").FindOne(ctx, bson.M{"username": username, "pin": pin}).Decode(&cm.info); err != nil {
		return false
	}

	return true
}

func (cm *CashMachine) WithdrawAdd(ctx context.Context, txType string) error {

	// User inputs
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

	// Abstracts the function to be able to withdraw and add
	switch txType {
	case "withdraw":
		cm.info.Amount -= amount
		break
	case "add":
		cm.info.Amount += amount
		break
	}

  fmt.Printf("New amount is £%g", cm.info.Amount)

	// Updates as such in the database

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

`)

	// User input for amount
	fmt.Println("Enter amount > ")
	var amount float64
	fmt.Scanln(&amount)

	if amount > cm.info.Amount {
		return amountTooBigErr
	}

	// ---- THIS BLOCK HANDLES RAND NUMBER GEN ----
	min := 0
	max := 100

	result := rand.Intn(max-min) + min

	fmt.Printf("You rolled %d", result)

	if result >= 51 {
		cm.info.Amount += amount * 2
	} else {
		cm.info.Amount -= amount / 2
	}

	fmt.Printf("\nNew amount £%g", cm.info.Amount)

	// ------------------------------------------

	// Updates as such in the database
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

	// Wages for each valid profession
	switch cm.info.Profession {
	case "developer":
		hourlyWage = 355
		workingHours = 8
	case "engineer":
		hourlyWage = 300
		workingHours = 7
	case "doctor":
		hourlyWage = 400
		workingHours = 10
	case "scientist":
		hourlyWage = 450
		workingHours = 10
	}

	newAmount := hourlyWage * workingHours
	cm.info.Amount += newAmount

	fmt.Printf("You earned £%g!", newAmount)

	// Updates as such in the database
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
