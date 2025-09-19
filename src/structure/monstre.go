package RED

import (
	RED "RED/menu"
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

func MenuAttaque(perso *Personnage) {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                ⚔️  ATTAQUE                    ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	for i, skill := range perso.Skill {
		fmt.Printf("║ [%d] %s                            ║\n", i+1, skill)
	}
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("👉 Choisissez un sort : ")
}

func GoblinPattern(gobelin *Monstre, perso *Personnage, tour int) {
	var degats int
	if tour%3 == 0 {
		degats = gobelin.Attaque * 2
		fmt.Printf("\n⚔️  %s prépare une attaque PUISSANTE !\n", gobelin.Nom)
	} else {
		degats = gobelin.Attaque
	}

	perso.PvActuels -= degats
	if perso.PvActuels < 0 {
		perso.PvActuels = 0
	}

	RED.Pause(2)
	fmt.Printf("⚔️ %s inflige à %s %d points de dégâts !\n", gobelin.Nom, perso.Nom, degats)
	fmt.Printf("💖 PV actuels : %d/%d\n", perso.PvActuels, perso.PvMax)
}

func CharacterTurn(perso *Personnage, gobelin *Monstre) bool {
	for {
		fmt.Println("\n╔═════════════════════════════════════════╗")
		fmt.Println("║               VOTRE TOUR                ║")
		fmt.Println("╠═════════════════════════════════════════╣")
		fmt.Println("║ [1] 📜 Infos du personnage              ║")
		fmt.Println("║ [2] ⚔️  Attaquer                         ║")
		fmt.Println("║ [3] 🎒 Inventaire                       ║")
		fmt.Println("║ [X] ❌ Fuir le combat                   ║")
		fmt.Println("╚═════════════════════════════════════════╝")
		fmt.Print("👉 Votre choix : ")

		choix := REDM.LireChoix()
		switch choix {
		case "1":
			REDM.ClearTerminal()
			DisplayInfo(*perso)
			fmt.Println("\nAppuyez sur Entrée pour revenir au tour...")
			fmt.Scanln()
			continue

		case "2":
			REDM.ClearTerminal()
			MenuAttaque(perso)
			sortChoisi := REDM.LireChoix()
			idx := -1
			fmt.Sscan(sortChoisi, &idx)

			var degats int = 5
			if idx >= 1 && idx <= len(perso.Skill) {
				skill := perso.Skill[idx-1]
				fmt.Printf("\n⚔️ %s utilise %s et inflige %d dégâts à %s !\n",
					perso.Nom, skill, degats, gobelin.Nom)
			} else {
				fmt.Printf("\n⚔️ %s utilise Attaque basique et inflige %d dégâts à %s !\n",
					perso.Nom, degats, gobelin.Nom)
			}
			gobelin.PvActuels -= degats
			if gobelin.PvActuels < 0 {
				gobelin.PvActuels = 0
			}
			fmt.Printf("💖 PV restants de %s : %d/%d\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)
			return true

		case "3":
			REDM.ClearTerminal()
			actionUtilisee := InventaireCombat(perso)
			if actionUtilisee {
				return true
			}
			continue

		case "X", "x":
			fmt.Println("💨 Vous avez essayé de fuir le combat MAIS...")
			RED.Pause(2)
			fmt.Println("👹 LE GOBELIN VOUS RATTRAPE !")
			RED.Pause(2)
			fmt.Println("FINISSEZ EN AVEC LUI UN BONNE FOIS POUR TOUTE !")
			RED.Pause(2)
			return false

		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

func TrainingFight(perso *Personnage) {
	gobelin := InitGobelin()
	tour := 1

	for gobelin.PvActuels > 0 && perso.PvActuels > 0 {
		REDM.ClearTerminal()
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Printf("║                   TOUR %d                     ║\n", tour)
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Printf("║                     PV                       ║\n")
		fmt.Println("╠══════════════════════════════════════════════╣")
		fmt.Printf("║💖 Vos PV : %d/%d                            ║\n", perso.PvActuels, perso.PvMax)
		fmt.Printf("║💖 PV du %s : %d/%d       ║\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)
		fmt.Println("╚══════════════════════════════════════════════╝")

		actionEffectuee := CharacterTurn(perso, &gobelin)

		if gobelin.PvActuels <= 0 {
			fmt.Println("🎉 Vous avez vaincu le gobelin !")
			REDM.Pause(2)
			fmt.Println("💰 Votre victoire vous a rapporté : 10 PO")
			perso.Gold += 10
			REDM.Pause(2)
			break
		}

		if actionEffectuee {
			GoblinPattern(&gobelin, perso, tour)

			if perso.PvActuels <= 0 {
				IsDead(perso)
				REDM.Pause(2)
			}
			tour++
		}

		fmt.Println("\nAppuyez sur Entrée pour passer au tour suivant...")
		fmt.Scanln()
	}
}
