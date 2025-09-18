package RED

import "fmt"

func AfficherMenu() {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║              🎮 MENU PRINCIPAL               ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ [1] 📜 Afficher les infos du personnage      ║")
	fmt.Println("║ [2] 🎒 Accéder à l'inventaire                ║")
	fmt.Println("║ [3] 🛒 Interface du Marchand                 ║")
	fmt.Println("║ [4] 📚 Sort du Joueur                        ║")
	fmt.Println("║ [5] ⚒️  Forgeron                              ║")
	fmt.Println("║ [A] ⚔️  Combattre un Gobelin d'entraînement   ║")
	fmt.Println("║ [X] ❌ Quitter                               ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix : ")
}

func LireChoix() string {
	var choix string
	fmt.Scanln(&choix)
	return choix
}

func MenuEnCombat() {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                🎮 MENU EN COMBAT             ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ [1] 📜 Informations du personnage           ║")
	fmt.Println("║ [2] 🎒 Accéder à l'inventaire               ║")
	fmt.Println("║ [3] 🔙 Retour au combat                     ║")
	fmt.Println("║ [X] ❌ Fuir le combat                       ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
}
