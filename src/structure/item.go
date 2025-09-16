package RED

import "fmt"

type Items struct {
	Items  []string
	Max    int
	Effet  []string
	Valeur int
}

func TakePotion(perso *Personnage) {
	for i, v := range perso.Inventaire {
		if v == "Potion" {
			if perso.PvActuels < perso.PvMax {
				perso.PvActuels += 50
				if perso.PvActuels > perso.PvMax {
					perso.PvActuels = perso.PvMax
				}
				Inv := []string{}
				for j, v := range perso.Inventaire {
					if j != i {
						Inv = append(Inv, v)
					}
				}
				perso.Inventaire = Inv
				fmt.Println("Vous avez utilisé une Potion. PV actuels :", perso.PvActuels)
				return
			} else {
				fmt.Println("Vos PV sont déjà au maximum.")
				return
			}
		}
	}
	fmt.Println("Vous n'avez pas de Potion dans votre inventaire.")
}
