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
		REDM.ClearTerminal()
		RED.DisplayInfo(*perso)
		choix2 := REDM.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "2":
		REDM.ClearTerminal()
		RED.InfoInventaire()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			REDM.ClearTerminal()
			RED.AccessInventory(*perso)
			choix3 := REDM.LireChoix()
			switch choix3 {
			case "1":
				Menu(perso)
			}
		case "2":
			REDM.ClearTerminal()
			RED.TakePot(perso)
			Menu(perso)
		case "3":
			REDM.ClearTerminal()
			RED.SpellBook(perso)
			Menu(perso)
		case "4":
			Menu(perso)
		}
	case "3":
		REDM.ClearTerminal()
		REDM.InterfaceMarchand()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Potion")
			Menu(perso)
		case "2":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Livre de Sort : Boule de Feu")
			Menu(perso)
		case "3":
			Menu(perso)
		}
	case "4":
		REDM.ClearTerminal()
		RED.InfoSort(perso)
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			Menu(perso)
		}
	case "5":
		REDM.ClearTerminal()
		fmt.Println("À bientôt !")
		return
	default:
		REDM.ClearTerminal()
		fmt.Println("Choix invalide.")
		Menu(perso)
	}
}
