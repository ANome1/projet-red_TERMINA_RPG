package main

import (
	RED "RED/src/structure"
	"fmt"
)

func main() {
	perso := RED.InitCharacter("Aragorn", "Ranger", 10, 100, 100, []string{"Arc", "Flèches", "Potion"})
	Menu(&perso)
}

func Menu(perso *RED.Personnage) {
	RED.AfficherMenu()
	choix := RED.LireChoix()

	switch choix {
	case "1":
		RED.DisplayInfo(*perso)
		choix2 := RED.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "2":
		RED.AccessInventory(*perso)
		choix2 := RED.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "3":
		RED.InterfaceMarchand()
		choix2 := RED.LireChoix()
		switch choix2 {
		case "1":
			RED.AddInventory(perso, "Potion")
			Menu(perso)
		case "2":
			Menu(perso)
		}
	case "4":
		fmt.Println("À bientôt !")
		return
	default:
		fmt.Println("Choix invalide.")
		Menu(perso)
	}
}
