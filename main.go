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

  _ = cm.Login(ctx)

}

func mainScreen() {
  
}

func authScreen() {
  var username string
  var pin int

  fmt.Println(UI_ART)

  fmt.Println(`
  - 1 - Create a new account
  - 2 - Login to existing account
  `)

  fmt.Println("> ")
  var option int
  fmt.Scanln(&option)

  if option == 1 {

  }

}

func main() {
	os.Setenv("MONGO_URI", "")

}
