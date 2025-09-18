package RED

import (
	REDM "RED/menu"
	"fmt"
)

type Monstre struct {
	Nom       string
	PvMax     int
	PvActuels int
	Attaque   int
}

func InitGobelin() Monstre {
	return Monstre{
		Nom:       "Gobelin d'entraînement",
		PvMax:     40,
		PvActuels: 40,
		Attaque:   5,
	}
}
func MenuGobelin(gobelin *Monstre, perso *Personnage) {
	REDM.ClearTerminal()
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║          ⚔️  COMBAT CONTRE UN GOBELIN         ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ Vous affrontez un Gobelin d'entraînement !   ║")
	fmt.Println("║ Montrez-lui de quoi vous êtes capable !      ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Printf("\n💖 Vous avez %d/%d points de vie.\n", perso.PvActuels, perso.PvMax)
	fmt.Printf("\n💖 Le %s a %d/%d points de vie.\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)
	REDM.Pause(3)
	REDM.ClearTerminal()
}

func GoblinPattern(gobelin *Monstre, perso *Personnage, tour int) {
	var degats int
	if tour%3 == 0 {
		degats = gobelin.Attaque * 2 // Double pour ce tour uniquement
	} else {
		degats = gobelin.Attaque
	}

	// Appliquer les dégâts
	perso.PvActuels -= degats
	if perso.PvActuels < 0 {
		perso.PvActuels = 0
	}

	// Affichage
	fmt.Printf("\n⚔️  %s inflige à %s %d points de dégâts !\n", gobelin.Nom, perso.Nom, degats)
	fmt.Printf("💖 PV actuels : %d/%d\n", perso.PvActuels, perso.PvMax)
}

func CharacterTurn(perso *Personnage, gobelin *Monstre, tour *int) {
	fmt.Println("\n╔═════════════════════════════════════════╗")
	fmt.Println("║               VOTRE TOUR                 ║")
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("║ [1] 📜 Menu                              ║")
	fmt.Println("║ [2] ⚔️ Attaquer                          ║")
	fmt.Println("║ [3] 🎒 Inventaire                        ║")
	fmt.Println("║ [X] ❌ Fuir le combat                    ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix 😊 ")

	choix := REDM.LireChoix()

	switch choix {
	case "1": // Menu principal
		REDM.MenuEnCombat() // appelle ton menu général
	case "2": // Attaquer
		degats := 5
		gobelin.PvActuels -= degats
		if gobelin.PvActuels < 0 {
			gobelin.PvActuels = 0
		}
		fmt.Printf("\n⚔️ %s utilise Attaque basique et inflige %d dégâts à %s !\n", perso.Nom, degats, gobelin.Nom)
		fmt.Printf("💖 PV restants de %s : %d/%d\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)

		if gobelin.PvActuels <= 0 {
			fmt.Println("🎉 Vous avez vaincu le gobelin !")
			return
		}

		// Tour du gobelin
		*tour++
		GoblinPattern(gobelin, perso, *tour)

	case "3": // Inventaire
		InventaireCombat(perso)
		choixInv := REDM.LireChoix()
		switch choixInv {
		case "1":
			TakePot(perso)
		case "2":
			fmt.Println("🛡️ Voir équipements non implémenté pour le combat")
		case "X", "x":
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}

	case "X", "x": // Fuir
		fmt.Println("🏃‍♂️ Vous avez fui le combat !")
		return

	default:
		fmt.Println("❌ Choix invalide.")
	}
}

func MenuAttaque(perso *Personnage) {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                ⚔️  ATTAQUE                    ║")
	fmt.Println("╠══════════════════════════════════════════════╣")

	for i, skill := range perso.Skill {
		fmt.Printf("║ [%d] %s\n", i+1, skill)
	}

	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix : ")
}
