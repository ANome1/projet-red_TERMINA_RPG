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
	case "1": // Informations personnage
		REDM.ClearTerminal()
		RED.DisplayInfo(*perso)
		choix2 := REDM.LireChoix()
		if choix2 == "1" {
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "2": // Inventaire
		REDM.ClearTerminal()
		RED.InfoInventaire()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1": // Inventaire général
			REDM.ClearTerminal()
			RED.AccessInventory(perso)
			choix3 := REDM.LireChoix()
			if choix3 == "1" {
				REDM.ClearTerminal()
				Menu(perso)
			}
		case "2": // Potions
			REDM.ClearTerminal()
			RED.InventairePotion(perso)
			choix5 := REDM.LireChoix()
			switch choix5 {
			case "1":
				REDM.ClearTerminal()
				RED.TakePot(perso)
				Menu(perso)
			case "2":
				REDM.ClearTerminal()
				RED.PoisonPot(perso)
				Menu(perso)
			case "3":
				REDM.ClearTerminal()
				Menu(perso)
			}
		case "3": // Livre de sorts
			REDM.ClearTerminal()
			RED.InventaireLivres(perso)
			choix3 := REDM.LireChoix()
			switch choix3 {
			case "1":
				REDM.ClearTerminal()
				RED.SpellBookFeu(perso)
				Menu(perso)
			case "2":
				REDM.ClearTerminal()
				RED.SpellBookInv(perso)
				RED.UpgradeInventory(perso)
				Menu(perso)
			case "3":
				REDM.ClearTerminal()
				Menu(perso)
			}
		case "4":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "3": // Marchand
		REDM.ClearTerminal()
		REDM.InterfaceMarchand()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			REDM.ClearTerminal()
			if perso.Gold >= 3 {
				RED.AddInventory(perso, "Potion de soin")
				perso.Gold -= 3
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "2":
			REDM.ClearTerminal()
			if perso.Gold >= 6 {
				RED.AddInventory(perso, "Potion de poison")
				perso.Gold -= 6
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "3":
			REDM.ClearTerminal()
			if perso.Gold >= 25 {
				RED.AddInventory(perso, "Livre de Sort : Boule de Feu")
				perso.Gold -= 25
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "4":
			REDM.ClearTerminal()
			if perso.Gold >= 30 {
				RED.AddInventory(perso, "Livre de Sort : UP Inventaire")
				perso.Gold -= 30
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "5":
			REDM.ClearTerminal()
			if perso.Gold >= 4 {
				RED.AddInventory(perso, "Fourrure de Loup")
				perso.Gold -= 4
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "6":
			REDM.ClearTerminal()
			if perso.Gold >= 7 {
				RED.AddInventory(perso, "Peau de Troll")
				perso.Gold -= 7
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "7":
			REDM.ClearTerminal()
			if perso.Gold >= 3 {
				RED.AddInventory(perso, "Cuir de Sanglier")
				perso.Gold -= 3
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "8":
			REDM.ClearTerminal()
			if perso.Gold >= 1 {
				RED.AddInventory(perso, "Plume de Corbeau")
				perso.Gold -= 1
			} else {
				fmt.Println("❌ Vous n'avez pas assez d'or pour acheter cet objet")
			}
			Menu(perso)
		case "9":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "4": // Sorts
		REDM.ClearTerminal()
		RED.InfoSort(perso)
		choix2 := REDM.LireChoix()
		if choix2 == "1" {
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "5": // Forgeron
		REDM.ClearTerminal()
		REDM.MenuForgeron()
		choixForge := REDM.LireChoix()
		switch choixForge {
		case "1":
			materiaux := map[string]int{"Plume de Corbeau": 1, "Cuir de Sanglier": 1}
			equip := RED.Equipements[0]
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "2":
			materiaux := map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}
			equip := RED.Equipements[1]
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "3":
			materiaux := map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}
			equip := RED.Equipements[2]
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "4":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "6": // Quitter
		REDM.ClearTerminal()
		fmt.Println("\n╔═════════════════════════════════════════════════╗")
		fmt.Println("║       🎮 A bientôt ! Merci d'avoir joué !       ║")
		fmt.Println("╚═════════════════════════════════════════════════╝")
		return

	default:
		REDM.ClearTerminal()
		fmt.Println("Choix invalide.")
		Menu(perso)
	}
}
