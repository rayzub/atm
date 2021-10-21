package main

import (
	"atm/src"
  tm "github.com/buger/goterm"
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
  // Initialise context that mongo requires, passed into each function as a param
	ctx := context.Background()
  // Initialise cash machine struct
	cm := &src.CashMachine{}

  // Connect to database
	cm.NewConnection(ctx)


  
  // Loops user input for easy error handling 
	for {
    fmt.Println(`
	- 1 - Create an account
	- 2 - Log in to existing account
`)
    fmt.Println("Choose an option: ")
		var option int
		fmt.Scanln(&option)

		if option == 1 {
			if created := cm.CreateAccount(ctx); created {
        continue
      }
		} else if option == 2 {
			if verified := cm.Login(ctx); verified {
        break
      }
    } else {
      continue
    }
	}

  tm.Clear()


  // Prints main screen after new account created or login works (as loop breaks)


  // Loops user input for easy error handling 
  for {
    fmt.Println(`
  - 1 - Withdraw
  - 2 - Add
  - 3 - Gamble
  - 4 - Work
  `)

    fmt.Println("Choose an option: ")
    var option int
    fmt.Scanln(&option)

    switch option {
      case 1:
        cm.WithdrawAdd(ctx, "withdraw")
        break
      case 2:
        cm.WithdrawAdd(ctx, "add")
        break
      case 3:
        cm.Gamble(ctx)
        break
      case 4:
        cm.Work(ctx)
        break
      default:
        continue
    }


  }
}
func main() {
	os.Setenv("MONGO_URI", )
	homeScreen()
}
