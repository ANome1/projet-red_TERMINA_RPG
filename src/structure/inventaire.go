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
}
