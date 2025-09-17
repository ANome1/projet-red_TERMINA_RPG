package RED

import (
	REDM "RED/menu"
	"fmt"
	"time"
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
func MenuGobelin(perso *Personnage) {
	REDM.ClearTerminal()
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║          ⚔️  COMBAT CONTRE UN GOBELIN         ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	fmt.Println("║ Vous affrontez un Gobelin d'entraînement !   ║")
	fmt.Println("║ Montrez-lui de quoi vous êtes capable !      ║")
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Printf("\n💖 Vous avez %d/%d points de vie.\n", perso.PvActuels, perso.PvMax)
	time.Sleep(3 * time.Second)
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
	for {
		// Menu principal du tour
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Println("║               ⚔️  VOTRE TOUR                  ║")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Println("║ [1] 📜 Menu                                   ║")
		fmt.Println("║ [2] ⚔️ Attaquer                                ║")
		fmt.Println("║ [3] 🎒 Inventaire                              ║")
		fmt.Println("║ [X] 🔙 Fuir le combat                          ║")
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Print("👉 Votre choix : ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			// Affiche le menu d’infos personnage
			DisplayInfo(*perso)
			continue

		case "2":
			// Menu d'attaque
			MenuAttaque(perso)

			var choixSkill int
			fmt.Scanln(&choixSkill)

			if choixSkill < 1 || choixSkill > len(perso.Skill) {
				fmt.Println("❌ Choix invalide, attaque annulée.")
				continue
			}

			// Dégâts appliqués (pour l'instant toutes les attaques infligent 5 dégâts)
			degats := 5
			gobelin.PvActuels -= degats
			if gobelin.PvActuels < 0 {
				gobelin.PvActuels = 0
			}

			fmt.Printf("\n💥 %s utilise %s et inflige %d points de dégâts à %s !\n", perso.Nom, perso.Skill[choixSkill-1], degats, gobelin.Nom)
			fmt.Printf("💖 PV restants du Gobelin : %d/%d\n", gobelin.PvActuels, gobelin.PvMax)
			*tour++

			// Vérification victoire
			if gobelin.PvActuels <= 0 {
				fmt.Printf("\n╔══════════════════════════════════════════════╗\n")
				fmt.Printf("║ 🎉 Vous avez vaincu %s !                          ║\n", gobelin.Nom)
				fmt.Println("╚══════════════════════════════════════════════╝")
				return
			}

			// Tour du Gobelin
			GoblinPattern(gobelin, perso, *tour)

			// Vérification défaite
			if perso.PvActuels <= 0 {
				fmt.Printf("\n╔══════════════════════════════════════════════╗\n")
				fmt.Printf("║ 💀 Vous avez été vaincu par %s...                 ║\n", gobelin.Nom)
				fmt.Println("╚══════════════════════════════════════════════╝")
				IsDead(perso)
				return
			}

			return

		case "3":
			// Inventaire
			InventaireCombat(perso)
			continue

		case "X", "x":
			fmt.Println("🔙 Vous fuyez le combat !")

			return

		default:
			fmt.Println("❌ Choix invalide, veuillez réessayer.")
		}
	}
}
