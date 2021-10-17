package main

import (
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

func main() {
	os.Setenv("MONGO_URI", "")
	mainScreen()
}

func mainScreen() {
	fmt.Println(UI_ART)
}

func authScreen() {}