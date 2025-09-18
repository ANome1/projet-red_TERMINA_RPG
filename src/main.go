package main

import (
	REDM "RED/menu"
	RED "RED/structure"
	"fmt"
)

func main() {
	REDM.ClearTerminal()
	fmt.Println("\n╔═════════════════════════════════════════════════╗")
	fmt.Println("║        🎮 BIENVENUE DANS TERMINA RPG            ║")
	fmt.Println("╚═════════════════════════════════════════════════╝")
	REDM.Pause(3)
	REDM.ClearTerminal()
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
		if choix2 == "x" || choix2 == "X" {
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
			if choix3 == "x" || choix3 == "X" {
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
			case "x", "X":
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
		case "4": // Équipements
			REDM.ClearTerminal()
			RED.InventaireEquipement(perso)
			fmt.Println("\n[1] 🛡️  Équiper un item")
			fmt.Println("[2] ❌ Déséquiper un item")
			fmt.Println("[x] 🔙 Retour au Menu Principal")
			choixEquip := REDM.LireChoix()

			switch choixEquip {
			case "1": // Équiper
				REDM.ClearTerminal() // Nettoie avant d’afficher les équipements disponibles
				fmt.Println("\nÉquipements disponibles à équiper :")
				for i, equip := range RED.Equipements {
					if RED.HasItem(perso, equip.Nom) {
						fmt.Printf("[%d] %s (%s)\n", i+1, equip.Nom, equip.Categorie)
					}
				}
				fmt.Print("Choisissez le numéro de l’équipement : ")
				num := REDM.LireChoix()

				REDM.ClearTerminal() // Nettoie avant d’afficher le résultat
				for i, equip := range RED.Equipements {
					if fmt.Sprint(i+1) == num {
						if RED.HasItem(perso, equip.Nom) {
							RED.Equiper(perso, equip)
							break
						} else {
							fmt.Println("❌ Vous ne possédez pas cet équipement")
						}
					}
				}

			case "2": // Déséquiper
				REDM.ClearTerminal() // Nettoie avant d’afficher les équipements équipés
				fmt.Println("\nÉquipements équipés :")
				if perso.Equipement.Tete != nil {
					fmt.Println("[1] Tête :", perso.Equipement.Tete.Nom)
				}
				if perso.Equipement.Torse != nil {
					fmt.Println("[2] Torse :", perso.Equipement.Torse.Nom)
				}
				if perso.Equipement.Pieds != nil {
					fmt.Println("[3] Pieds :", perso.Equipement.Pieds.Nom)
				}
				fmt.Print("Choisissez le numéro de la catégorie à déséquiper : ")
				num := REDM.LireChoix()

				REDM.ClearTerminal() // Nettoie avant d’afficher le résultat
				switch num {
				case "1":
					RED.Desequiper(perso, "Tête")
				case "2":
					RED.Desequiper(perso, "Torse")
				case "3":
					RED.Desequiper(perso, "Pieds")
				default:
					fmt.Println("❌ Choix invalide")
				}

			case "x", "X":
				REDM.ClearTerminal()
			}
			Menu(perso)
		case "x", "X":
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
		case "x", "X":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "4": // Sorts
		REDM.ClearTerminal()
		RED.InfoSort(perso)
		choix2 := REDM.LireChoix()
		if choix2 == "x" || choix2 == "X" {
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
			REDM.ClearTerminal()
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "2":
			materiaux := map[string]int{"Fourrure de Loup": 2, "Peau de Troll": 1}
			equip := RED.Equipements[1]
			REDM.ClearTerminal()
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "3":
			materiaux := map[string]int{"Fourrure de Loup": 1, "Cuir de Sanglier": 1}
			equip := RED.Equipements[2]
			REDM.ClearTerminal()
			RED.Forger(perso, equip, materiaux)
			Menu(perso)
		case "x", "X":
			REDM.ClearTerminal()
			Menu(perso)
		}

	case "A", "a": // Combat contre un gobelin d'entrainement
		REDM.ClearTerminal()
		RED.TrainingFight(perso) // Lancement du combat
		Menu(perso)              // Retour au menu principal après le combat

	case "x", "X": // Quitter
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
