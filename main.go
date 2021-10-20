package main

import (
	"atm/src"
	"context"
	"fmt"
	"os"
)

const UI_ART = `
 _______  _______                    _______  _          _______ _________ _______ 
(  ____ )(  ___  )|\     /||\     /|(  ___  )( (    /|  (  ___  )\__   __/(       )
| (    )|| (   ) |( \   / )| )   ( || (   ) ||  \  ( |  | (   ) |   ) (   | () () |
| (____)|| (___) | \ (_) / | (___) || (___) ||   \ | |  | (___) |   | |   | || || |
|     __)|  ___  |  \   /  |  ___  ||  ___  || (\ \) |  |  ___  |   | |   | |(_)| |
| (\ (   | (   ) |   ) (   | (   ) || (   ) || | \   |  | (   ) |   | |   | |   | |
| ) \ \__| )   ( |   | |   | )   ( || )   ( || )  \  |  | )   ( |   | |   | )   ( |
|/   \__/|/     \|   \_/   |/     \||/     \||/    )_)  |/     \|   )_(   |/     \|

`

func homeScreen() {
	fmt.Println(UI_ART)
	ctx := context.Background()
	cm := &src.CashMachine{}

	cm.NewConnection(ctx)

	fmt.Println(`
	- 1 - Create an account
	- 2 - Log in to existing account
`)
	for {
		fmt.Println("Choose an option: ")
		var option int
		fmt.Scanln(&option)

		if option == 1 {
			created := cm.CreateAccount(ctx)
			if created {
				break
			}
		} else if option == 2 {
			verified := cm.Login(ctx)
			if verified {
				break
			}
		} else {
			continue
		}
	}

}

func mainScreen() {

}

func authScreen() {

	fmt.Println(UI_ART)

	fmt.Println(`
  - 1 - Create a new account
  - 2 - Login to existing account
  `)

}

func main() {
	os.Setenv("MONGO_URI", "")
	homeScreen()
}
