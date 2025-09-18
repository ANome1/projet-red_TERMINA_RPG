package RED

import (
	REDM "RED/menu"
	"fmt"
)

// Structure du monstre
type Monstre struct {
	Nom       string
	PvMax     int
	PvActuels int
	Attaque   int
}

// Initialisation d'un Gobelin d'entraînement
func InitGobelin() Monstre {
	return Monstre{
		Nom:       "Gobelin d'entraînement",
		PvMax:     40,
		PvActuels: 40,
		Attaque:   5,
	}
}

// Affichage du menu de combat et des PV
func MenuGobelin(gobelin *Monstre, perso *Personnage) {
	REDM.ClearTerminal()
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║          ⚔️  COMBAT CONTRE UN GOBELIN         ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ Vous affrontez un Gobelin d'entraînement !   ║")
	fmt.Println("║ Montrez-lui de quoi vous êtes capable !      ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Printf("\n💖 Vous avez %d/%d points de vie.\n", perso.PvActuels, perso.PvMax)
	fmt.Printf("💖 Le %s a %d/%d points de vie.\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)
	REDM.Pause(3)
	REDM.ClearTerminal()
}

// Pattern d'attaque du Gobelin
func GoblinPattern(gobelin *Monstre, perso *Personnage, tour int) {
	degats := gobelin.Attaque
	if tour%3 == 0 {
		degats *= 2
	}

	perso.PvActuels -= degats
	if perso.PvActuels < 0 {
		perso.PvActuels = 0
	}

	fmt.Printf("\n⚔️ %s inflige à %s %d points de dégâts !\n", gobelin.Nom, perso.Nom, degats)
	fmt.Printf("💖 PV actuels : %d/%d\n", perso.PvActuels, perso.PvMax)

	if perso.PvActuels == 0 {
		IsDead(perso)
	}
}

// Menu des sorts du joueur
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

// Tour du joueur
func CharacterTurn(perso *Personnage, gobelin *Monstre, tour *int) {
	fmt.Println("\n╔═════════════════════════════════════════╗")
	fmt.Println("║               VOTRE TOUR                 ║")
	fmt.Println("╠═════════════════════════════════════════╣")
	fmt.Println("║ [1] 📜 Menu                              ║")
	fmt.Println("║ [2] ⚔️ Attaquer                          ║")
	fmt.Println("║ [3] 🎒 Inventaire                        ║")
	fmt.Println("║ [X] ❌ Fuir le combat                    ║")
	fmt.Println("╚═════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix : ")

	choix := REDM.LireChoix()

	switch choix {
	case "1": // Menu principal
		REDM.MenuEnCombat()
	case "2": // Attaquer
		MenuAttaque(perso) // Affiche les sorts disponibles
		fmt.Print("👉 Choisissez un sort : ")
		sortChoisi := REDM.LireChoix()
		var degats int
		idx := -1
		fmt.Sscan(sortChoisi, &idx)
		if idx >= 1 && idx <= len(perso.Skill) {
			skill := perso.Skill[idx-1]
			degats = 5 // par défaut pour tous les sorts ici
			fmt.Printf("\n⚔️ %s utilise %s et inflige %d dégâts à %s !\n",
				perso.Nom, skill, degats, gobelin.Nom)
			gobelin.PvActuels -= degats
			if gobelin.PvActuels < 0 {
				gobelin.PvActuels = 0
			}
			fmt.Printf("💖 PV restants de %s : %d/%d\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)
		} else {
			fmt.Println("❌ Choix invalide, attaque basique utilisée !")
			degats = 5
			gobelin.PvActuels -= degats
			if gobelin.PvActuels < 0 {
				gobelin.PvActuels = 0
			}
			fmt.Printf("⚔️ %s inflige %d dégâts à %s !\n", perso.Nom, degats, gobelin.Nom)
		}

		if gobelin.PvActuels <= 0 {
			fmt.Println("🎉 Vous avez vaincu le gobelin !")
			return
		}

		// Tour du gobelin
		*tour++
		GoblinPattern(gobelin, perso, *tour)

	case "3": // Inventaire
		InventaireCombat(perso)
		// Après avoir utilisé l'inventaire, c'est au tour du gobelin
		*tour++
		GoblinPattern(gobelin, perso, *tour)

	case "X", "x": // Fuir
		fmt.Println("🏃‍♂️ Vous avez fui le combat !")
		return

	default:
		fmt.Println("❌ Choix invalide.")
	}
}

func TrainingFight(perso *Personnage) {
	// Initialisation du gobelin
	gobelin := InitGobelin()

	// Affichage du menu de début de combat
	MenuGobelin(&gobelin, perso)

	// Tour de combat
	tour := 1

	// Boucle de combat
	for perso.PvActuels > 0 && gobelin.PvActuels > 0 {
		fmt.Printf("\n--- Tour %d ---\n", tour)
		REDM.Pause(2)

		// Tour du joueur
		CharacterTurn(perso, &gobelin, &tour)

		// Vérifier si le gobelin est vaincu
		if gobelin.PvActuels <= 0 {
			fmt.Println("🎉 Félicitations ! Vous avez gagné le combat d'entraînement !")
			break
		}

		// Tour du gobelin
		fmt.Println("\n--- Tour du Gobelin ---")
		REDM.Pause(2)
		GoblinPattern(&gobelin, perso, tour)

		// Vérifier si le joueur est mort
		if perso.PvActuels <= 0 {
			fmt.Println("💀 Vous avez été vaincu par le gobelin...")
			IsDead(perso)
			break
		}

		// Affichage des PV à la fin du tour
		fmt.Printf("💖 PV du joueur : %d/%d | PV du gobelin : %d/%d\n",
			perso.PvActuels, perso.PvMax, gobelin.PvActuels, gobelin.PvMax)

		// Incrémentation du tour
		tour++
	}

	fmt.Println("🏁 Combat terminé !")
}
