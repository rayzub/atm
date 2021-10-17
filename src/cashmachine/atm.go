package cashmachine

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type CashMachine struct {
	amount float64
	client *mongo.Client
}

type ATM interface {
	Withdraw()
	Login() bool
	Add()
	Work()
	Gamble()
}

func (cm *CashMachine) NewConnection(ctx context.Context) {

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal(err)
	}

	cm.client = client
}

func (cm *CashMachine) Withdraw() {}
func (cm *CashMachine) Add()      {}
func (cm *CashMachine) Gamble()   {}
func (cm *CashMachine) Work()     {}
func (cm *CashMachine) Login()    {}
