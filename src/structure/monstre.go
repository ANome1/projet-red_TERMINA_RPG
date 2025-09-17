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
	fmt.Println("\n╔═════════════════════════════════════════╗")
	fmt.Println("║               VOTRE TOUR                 ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║ [1] 📜 Menu                              ║")
	fmt.Println("║ [2] ⚔️ Attaquer                          ║")
	fmt.Println("║ [3] 🎒 Inventaire                        ║")
	fmt.Println("╚══════════════════════════════════════════╝")
	fmt.Print("👉 Votre choix 😊 ")
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
