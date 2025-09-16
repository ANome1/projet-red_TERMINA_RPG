package main

import (
	REDMenu "RED/menu"
	REDStruct "RED/structure"

	"fmt"
)

func main() {
	perso := REDStruct.InitCharacter("Aragorn", "Ranger", 10, 100, 100, []string{"Arc", "Flèches", "Potion"})
	Menu(&perso)
}

func Menu(perso *REDStruct.Personnage) {
	REDMenu.AfficherMenu()
	choix := REDMenu.LireChoix()

	switch choix {
	case "1":
		REDStruct.DisplayInfo(*perso)
		choix2 := REDMenu.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "2":
		REDStruct.AccessInventory(*perso)
		choix2 := REDMenu.LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "3":
		REDMenu.InterfaceMarchand()
		choix2 := REDMenu.LireChoix()
		switch choix2 {
		case "1":
			REDStruct.AddInventory(perso, "Potion")
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
