package main

import (
	REDM "RED/menu"
	RED "RED/structure"
	"fmt"
)

func main() {
	perso := RED.CharacterCreation()
	Menu(&perso)
}

func Menu(perso *RED.Personnage) {
	REDM.AfficherMenu()
	choix := REDM.LireChoix()

	switch choix {
	case "1":
		RED.DisplayInfo(*perso)
		choix2 := REDM.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "2":
		RED.InfoInventaire()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			RED.AccessInventory(*perso)
			choix3 := REDM.LireChoix()
			switch choix3 {
			case "1":
				Menu(perso)
			}
		case "3":
			Menu(perso)
		}
	case "3":
		REDM.InterfaceMarchand()
		choix2 := REDM.LireChoix()
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
