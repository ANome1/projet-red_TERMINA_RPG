package RED

import "fmt"

type Inventaire struct {
	Items []string
	Max   int
}

func AccessInventory(perso Personnage) {
	fmt.Println("=== Inventaire Joueur ===")
	for _, c := range perso.Inventaire {
		fmt.Println(c)
	}
	fmt.Println("1. Retour au Menu Principal")
}

func AddInventory(perso *Personnage, item string) {
	if InventairePlein(perso) == false {
		perso.Inventaire = append(perso.Inventaire, item)
		fmt.Println("Vous avez obtenu :", item)
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
	}
}

func RemoveInventory(perso *Personnage, item string) {
	for i, v := range perso.Inventaire {
		if v == item {
			Inv := []string{}
			for j, v := range perso.Inventaire {
				if j != i {
					Inv = append(Inv, v)
				}
			}
			perso.Inventaire = Inv
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println(item, "n'est pas dans l'inventaire.")
}

func InventairePlein(perso *Personnage) bool {
	return len(perso.Inventaire) >= 10
}
