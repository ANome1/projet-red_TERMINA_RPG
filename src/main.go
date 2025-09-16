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
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "2":
		REDM.ClearTerminal()
		RED.InfoInventaire()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			REDM.ClearTerminal()
			RED.AccessInventory(perso)
			choix3 := REDM.LireChoix()
			if choix3 == "1" {
				REDM.ClearTerminal()
				Menu(perso)
			}
		case "2":
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
		case "3":
			REDM.ClearTerminal()
			RED.SpellBook(perso)
			Menu(perso)
		case "4":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "3":
		REDM.ClearTerminal()
		REDM.InterfaceMarchand()
		choix2 := REDM.LireChoix()
		switch choix2 {
		case "1":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Potion de soin")
			perso.Gold -= 3
			Menu(perso)
		case "2":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Potion de poison")
			perso.Gold -= 6
			Menu(perso)
		case "3":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Livre de Sort : Boule de Feu")
			perso.Gold -= 25
			Menu(perso)
		case "4":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Fourrure de Loup")
			perso.Gold -= 4
			Menu(perso)
		case "5":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Peau de Troll")
			perso.Gold -= 7
			Menu(perso)
		case "6":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Cuir de Sanglier")
			perso.Gold -= 3
			Menu(perso)
		case "7":
			REDM.ClearTerminal()
			RED.AddInventory(perso, "Plume de Corbeau")
			perso.Gold -= 1
			Menu(perso)
		case "8":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "4":
		REDM.ClearTerminal()
		RED.InfoSort(perso)
		choix2 := REDM.LireChoix()
		if choix2 == "1" {
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "5":
		REDM.ClearTerminal()
		REDM.MenuForgeron()
		choixForge := REDM.LireChoix()
		switch choixForge {
		case "1":
			REDM.ClearTerminal()
			if RED.HasItem(perso, "Plume de Corbeau") && RED.HasItem(perso, "Cuir de Sanglier") {
				RED.RemoveInventory(perso, "Plume de Corbeau")
				RED.RemoveInventory(perso, "Cuir de Sanglier")
				RED.AddInventory(perso, "Chapeau de l'aventurier")
				perso.Gold -= 5
			} else {
				fmt.Println("❌ Vous n'avez pas les matériaux nécessaires pour forger le Chapeau de l'aventurier.")
			}
			Menu(perso)

		case "2":
			REDM.ClearTerminal()
			if RED.CountItem(perso, "Fourrure de Loup") >= 2 && RED.HasItem(perso, "Peau de Troll") {
				RED.RemoveInventory(perso, "Fourrure de Loup")
				RED.RemoveInventory(perso, "Fourrure de Loup")
				RED.RemoveInventory(perso, "Peau de Troll")
				RED.AddInventory(perso, "Tunique de l'aventurier")
				perso.Gold -= 5
			} else {
				fmt.Println("❌ Vous n'avez pas les matériaux nécessaires pour forger la Tunique de l'aventurier.")
			}
			Menu(perso)

		case "3":
			REDM.ClearTerminal()
			if RED.HasItem(perso, "Fourrure de Loup") && RED.HasItem(perso, "Cuir de Sanglier") {
				RED.RemoveInventory(perso, "Fourrure de Loup")
				RED.RemoveInventory(perso, "Cuir de Sanglier")
				RED.AddInventory(perso, "Bottes de l'aventurier")
				perso.Gold -= 5
			} else {
				fmt.Println("❌ Vous n'avez pas les matériaux nécessaires pour forger les Bottes de l'aventurier.")
			}
			Menu(perso)

		case "4":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "6":
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
