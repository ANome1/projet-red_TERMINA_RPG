package RED

import (
	"fmt"
	"time"
)

type Items struct {
	Items  []string
	Max    int
	Effet  []string
	Valeur int
}

func PoisonPot(perso *Personnage) {
	// Vérifie si le joueur a une Potion de Poison
	aPotion := false
	for _, item := range perso.Inventaire {
		if item == "Potion de poison" {
			aPotion = true
			break
		}
	}

	if !aPotion {
		fmt.Println("❌ Vous n'avez pas de Potion de Poison dans votre inventaire.")
		return
	}

	// Retire la potion avant d'appliquer l'effet
	RemoveInventory(perso, "Potion de Poison")

	fmt.Println("\n╔═════════════════════════════════════════════════╗")
	fmt.Println("║ ☠️  Vous avez utilisé une Potion de Poison !     ║")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		perso.PvActuels -= 10
		if perso.PvActuels < 0 {
			perso.PvActuels = 0
		}
		fmt.Printf("║ 💀 Dégâts empoisonnés :%-25d║\n", 10*i)
		fmt.Printf("║ ❤️  PV actuels :%-28d     ║\n", perso.PvActuels)
	}
	fmt.Println("╚═════════════════════════════════════════════════╝")

	// Vérifie si le joueur meurt
	if perso.PvActuels == 0 {
		IsDead(perso)
	}
}

func TakePot(perso *Personnage) {
	for i, v := range perso.Inventaire {
		if v == "Potion de soin" {
			if perso.PvActuels < perso.PvMax {
				perso.PvActuels += 50
				if perso.PvActuels > perso.PvMax {
					perso.PvActuels = perso.PvMax
				}
				Inv := []string{}
				for j, item := range perso.Inventaire {
					if j != i {
						Inv = append(Inv, item)
					}
				}
				perso.Inventaire = Inv

				fmt.Println("\n╔════════════════════════════════════════════╗")
				fmt.Println("║ 🧪 Vous avez utilisé une Potion !          ║")
				fmt.Printf("║ ❤️  PV actuels :%-28d║\n", perso.PvActuels)
				fmt.Println("╚════════════════════════════════════════════╝")
				return
			} else {
				fmt.Println("✅ Vos PV sont déjà au maximum.")
			}
		}
	}
	fmt.Println("❌ Vous n'avez pas de Potion dans votre inventaire.")
}
