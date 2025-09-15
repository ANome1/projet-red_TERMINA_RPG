package main

import (
	RED "RED/src/structure"
	"fmt"
)

func main() {
	perso := RED.InitCharacter("Aragorn", "Ranger", 10, 100, 100, []string{"Arc", "Flèches", "Potion"})
	for {
		RED.AfficherMenu()
		choix := RED.LireChoix()
		switch choix {
		case "1":
			RED.DisplayInfo(perso)
		case "2":
			RED.AccessInventory(perso)
		case "3":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
