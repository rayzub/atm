package main

import (
	"atm/src/cashmachine"
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

func mainScreen() {
	fmt.Println(UI_ART)
	ctx := context.Background()
	cm := &cashmachine.CashMachine{}

	cm.NewConnection(ctx)

	verified := cm.Login(ctx)
	if !verified {
		return
	}

}

func authScreen() {}

func main() {
	os.Setenv("MONGO_URI", "")

}
