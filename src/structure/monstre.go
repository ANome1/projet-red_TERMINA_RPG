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

// Menu des sorts
func MenuAttaque(perso *Personnage) {
	fmt.Println("\n╔══════════════════════════════════════════════╗")
	fmt.Println("║                ⚔️  ATTAQUE                    ║")
	fmt.Println("╠══════════════════════════════════════════════╣")
	for i, skill := range perso.Skill {
		fmt.Printf("║ [%d] %s\n", i+1, skill)
	}
	fmt.Println("╚══════════════════════════════════════════════╝")
	fmt.Print("👉 Choisissez un sort : ")
}

// Pattern d'attaque du Gobelin
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

	fmt.Printf("⚔️ %s inflige à %s %d points de dégâts !\n", gobelin.Nom, perso.Nom, degats)
	fmt.Printf("💖 PV actuels : %d/%d\n", perso.PvActuels, perso.PvMax)
}

// Tour du joueur
// Retourne true si une action concrète a été effectuée
func CharacterTurn(perso *Personnage, gobelin *Monstre) bool {
	for {
		fmt.Println("\n╔═════════════════════════════════════════╗")
		fmt.Println("║               VOTRE TOUR                 ║")
		fmt.Println("╠═════════════════════════════════════════╣")
		fmt.Println("║ [1] 📜 Infos du personnage               ║")
		fmt.Println("║ [2] ⚔️ Attaquer                          ║")
		fmt.Println("║ [3] 🎒 Inventaire                        ║")
		fmt.Println("║ [X] ❌ Fuir le combat                    ║")
		fmt.Println("╚═════════════════════════════════════════╝")
		fmt.Print("👉 Votre choix : ")

		choix := REDM.LireChoix()
		switch choix {
		case "1": // Afficher infos joueur
			REDM.ClearTerminal()
			DisplayInfo(*perso)
			fmt.Println("\nAppuyez sur Entrée pour revenir au tour...")
			fmt.Scanln()
			// Ne consomme pas le tour, reste dans la boucle
			continue

		case "2": // Attaquer
			MenuAttaque(perso) // Affiche les sorts disponibles
			fmt.Print("👉 Choisissez un sort : ")
			sortChoisi := REDM.LireChoix()
			idx := -1
			fmt.Sscan(sortChoisi, &idx)

			var degats int = 5 // Tous les sorts font 5 dégâts pour l'instant
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
			return true // Action effectuée, fin du tour joueur

		case "3": // Inventaire
			for {
				fmt.Println("\n╔════════════════════════════════════════════╗")
				fmt.Println("║           🎒 INVENTAIRE DU JOUEUR          ║")
				fmt.Println("╠════════════════════════════════════════════╣")
				fmt.Println("║ [1] 🧪 Utiliser une potion                 ║")
				fmt.Println("║ [2] 🛡️ Liste d'équipements                 ║")
				fmt.Println("║ [X] 🔙 Retour au Menu de combat            ║")
				fmt.Println("╚════════════════════════════════════════════╝")
				fmt.Print("👉 Votre choix : ")
				choixInv := REDM.LireChoix()

				if choixInv == "1" {
					InventairePotion(perso) // Affiche les potions
					fmt.Print("👉 Choisissez la potion : ")
					choixPot := REDM.LireChoix()
					switch choixPot {
					case "1":
						TakePot(perso)
						return true // action consommée
					case "2":
						PoisonPot(perso)
						return true // action consommée
					case "X", "x":
						continue // reste dans inventaire
					default:
						fmt.Println("❌ Choix invalide.")
					}
				} else if choixInv == "2" {
					fmt.Println("🛡️ Voir équipements en combat :")
					InventaireEquipement(perso)
					fmt.Println("[X] 🔙 Retour")
					fmt.Print("👉 Choisissez l'équipement : ")
					choixEquip := REDM.LireChoix()
					if choixEquip == "X" || choixEquip == "x" {
						continue // reste dans inventaire
					}
					// Gestion équipement si nécessaire
					// Si tu veux équiper/déséquiper, ici tu peux implémenter
					// Pour l'instant, on considère que regarder ne consomme pas le tour
				} else if choixInv == "X" || choixInv == "x" {
					break // Retour au menu principal
				} else {
					fmt.Println("❌ Choix invalide.")
				}
			}
			// Ne consomme pas le tour si on n'a rien utilisé
			continue

		case "X", "x": // Fuir
			fmt.Println("🏃‍♂️ Vous avez fui le combat !")
			return false // pas d'action concrète mais on quitte le combat

		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

// Combat d'entraînement
func TrainingFight(perso *Personnage) {
	gobelin := InitGobelin()
	tour := 1

	for gobelin.PvActuels > 0 && perso.PvActuels > 0 {
		REDM.ClearTerminal()
		fmt.Println("\n╔══════════════════════════════════════════════╗")
		fmt.Printf("║                 TOUR %d                        ║\n", tour)
		fmt.Println("╚══════════════════════════════════════════════╝")
		fmt.Printf("\n💖 Vos PV : %d/%d\n", perso.PvActuels, perso.PvMax)
		fmt.Printf("💖 PV du %s : %d/%d\n", gobelin.Nom, gobelin.PvActuels, gobelin.PvMax)

		actionEffectuee := CharacterTurn(perso, &gobelin)

		if gobelin.PvActuels <= 0 {
			fmt.Println("🎉 Vous avez vaincu le gobelin !")
			REDM.Pause(2)
			break
		}

		if actionEffectuee {
			// Tour du gobelin seulement si le joueur a agi
			GoblinPattern(&gobelin, perso, tour)

			if perso.PvActuels <= 0 {
				IsDead(perso)
				fmt.Println("💀 Vous avez été ressuscité avec 50% de vos PV max !")
				REDM.Pause(2)
			}
			tour++
		}

		fmt.Println("\nAppuyez sur Entrée pour passer au tour suivant...")
		fmt.Scanln()
	}
}
