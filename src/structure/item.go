package RED

import "fmt"

type Items struct {
	Items  []string
	Max    int
	Effet  []string
	Valeur int
}

func TakePot(perso *Personnage) {
	for i, v := range perso.Inventaire {
		if v == "Potion" {
			if perso.PvActuels < perso.PvMax {
				perso.PvActuels += 50
				if perso.PvActuels > perso.PvMax {
					perso.PvActuels = perso.PvMax
				}

				// Retirer la potion de l'inventaire
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
				return
			}
		}
	}
	fmt.Println("❌ Vous n'avez pas de Potion dans votre inventaire.")
}
