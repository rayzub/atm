package src

import (
	"context"
  "time"
  "errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)


var (
  amountTooBigErr = errors.New("Amount is larger than amount in bank")
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

func (cm *CashMachine) WithdrawAdd(ctx context.Context, txType string) error {

  fmt.Println("Enter amount to withdraw > ")
  var amount float64
  _, err := fmt.Scanln(&amount); if err != nil {
    log.Fatalf("ATM - Withdaw - Error %v", err); return err  
  }

  if amount > cm.info.Amount {
    return amountTooBigErr
  }


  switch txType {
    case "withdraw":
      cm.info.Amount - amount
      break
    case "add":
      cm.info.Amount + amount
      break
  }

  tx := Transaction{
    Type: txType,
    Amount: amount,
    User: os.Getenv("USERNAME"),
    Date: time.Now().Unix(),
  }

  _, err := cm.client.Collection("transactions").InsertOne(ctx, tx)

  if err != nil {
    log.Fatalf("ATM - Withdraw - Database error %v", err); return err
  }
  
  return nil
}
func (cm *CashMachine) Gamble(ctx context.Context)   {}
func (cm *CashMachine) Work(ctx context.Context)     {}
