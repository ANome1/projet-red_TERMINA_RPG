package RED

import "fmt"

type Inventaire struct {
	Items []string
	Max   int
}

func InventairePotion(perso *Personnage) {
	potionsDisponibles := []string{"Potion de soin", "Potion de poison"}
	compteur := make(map[string]int)
	for _, item := range perso.Inventaire {
		compteur[item]++
	}
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║           🎒 LISTE DES POTIONS             ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	for _, potion := range potionsDisponibles {
		qte := compteur[potion]
		fmt.Printf("║ • %-28s x%-3d        ║\n", potion, qte)
	}
	fmt.Println("║ 🔙 Retour au Menu Principal [3]            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func AccessInventory(perso *Personnage) {
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║           🎒 INVENTAIRE DU JOUEUR          ║")
	fmt.Println("╠════════════════════════════════════════════╣")

	if len(perso.Inventaire) == 0 {
		fmt.Println("║ (Inventaire vide)                          ║")
	} else {
		for _, item := range perso.Inventaire {
			if len(item) > 36 {
				item = item[:33] + "..."
			}
			fmt.Printf("║ • %-38s   ║\n", item)
		}
	}
	fmt.Println("║ 🔙 Retour au Menu Principal [1]            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func InfoInventaire() {
	fmt.Println("\n╔═════════════════════════════════╗")
	fmt.Println("║     INFORMATIONS INVENTAIRE     ║")
	fmt.Println("╠═════════════════════════════════╣")
	fmt.Println("║ [1] 👀 Afficher l'inventaire    ║")
	fmt.Println("║ [2] 🧪 Utiliser une potion      ║ ")
	fmt.Println("║ [3] 📚 Utiliser un livre de sort║ ")
	fmt.Println("║ [4] 🔙 Retour au Menu Principal ║")
	fmt.Println("╚═════════════════════════════════╝")
	fmt.Print("👉 Votre choix : ")
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
	fmt.Println(item, "n'a pas été trouvé dans l'inventaire.")
}

func InventairePlein(perso *Personnage) bool {
	return len(perso.Inventaire) >= 10
}

func HasItem(perso *Personnage, item string) bool {
	for _, i := range perso.Inventaire {
		if i == item {
			return true
		}
	}
	return false
}

func CountItem(perso *Personnage, item string) int {
	count := 0
	for _, i := range perso.Inventaire {
		if i == item {
			count++
		}
	}
	return count
}
