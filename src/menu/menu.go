package RED

import (
	RED "RED/structure"
	"fmt"
)

func Menu(perso *RED.Personnage) {
	AfficherMenu()
	choix := LireChoix()

	switch choix {
	case "1":
		RED.DisplayInfo(*perso)
		choix2 := LireChoix()
		if choix2 == "1" {
			Menu(perso)
		}
	case "2":
		RED.InfoInventaire()
		choix2 := LireChoix()
		switch choix2 {
		case "1":
			RED.AccessInventory(*perso)
			choix3 := LireChoix()
			switch choix3 {
			case "1":
				Menu(perso)
			}
		case "2":
			RED.TakePot(perso)
		case "3":
			Menu(perso)
		}
	case "3":
		InterfaceMarchand()
		choix2 := LireChoix()
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

func AfficherMenu() {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║              🎮 MENU PRINCIPAL               ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ 1. 📜 Afficher les infos du personnage       ║")
	fmt.Println("║ 2. 🎒 Accéder à l'inventaire                 ║")
	fmt.Println("║ 3. 🛒 Interface du Marchand                  ║")
	fmt.Println("║ 4. ❌ Quitter                                ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix : ")
}

func LireChoix() string {
	var choix string
	fmt.Scanln(&choix)
	return choix
}
