package RED

import "fmt"

type Inventaire struct {
	Items      []string
	Max        int
	CptUpgrade int
}

func InventairePotion(perso *Personnage) {
	potionsDisponibles := []string{"Potion de soin", "Potion de poison"}
	compteur := make(map[string]int)
	for _, item := range perso.Inventaire.Items {
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

func InventaireLivres(perso *Personnage) {
	livresDisponibles := []string{
		"Livre de Sort : Boule de Feu",
		"Livre de Sort : UP Inventaire",
	}
	compteur := make(map[string]int)
	for _, item := range perso.Inventaire.Items {
		compteur[item]++
	}

	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║            📚 LIVRES DE SORT               ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	for _, livre := range livresDisponibles {
		qte := compteur[livre]
		fmt.Printf("║ • %-36s x%-3d║\n", livre, qte)
	}
	fmt.Println("║ 🔙 Retour au Menu Principal [3]            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func AccessInventory(perso *Personnage) {
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║           🎒 INVENTAIRE DU JOUEUR          ║")
	fmt.Println("╠════════════════════════════════════════════╣")

	if len(perso.Inventaire.Items) == 0 {
		fmt.Println("║ (Inventaire vide)                          ║")
	} else {
		for _, item := range perso.Inventaire.Items {
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
	if !InventairePlein(perso) {
		perso.Inventaire.Items = append(perso.Inventaire.Items, item)
		fmt.Println("Vous avez obtenu :", item)
	} else {
		fmt.Println("Inventaire plein ! Impossible d'ajouter", item)
	}
}

func RemoveInventory(perso *Personnage, item string) {
	for i, v := range perso.Inventaire.Items {
		if v == item {
			Inv := []string{}
			for j, v := range perso.Inventaire.Items {
				if j != i {
					Inv = append(Inv, v)
				}
			}
			perso.Inventaire.Items = Inv
			fmt.Println(item, "a été retiré de l'inventaire.")
			return
		}
	}
	fmt.Println(item, "n'a pas été trouvé dans l'inventaire.")
}

func InventairePlein(perso *Personnage) bool {
	return len(perso.Inventaire.Items) >= 10
}

func UpgradeInventory(perso *Personnage) {
	if perso.Inventaire.CptUpgrade >= 3 {
		fmt.Println("❌ Vous avez atteint votre limite d'augmentation d'inventaire")
	} else {
		perso.Inventaire.Max += 10
		perso.Inventaire.CptUpgrade += 1
		fmt.Println("✅ Félicitations ! La taille de votre inentaire a été augmenté de : 10")
	}
}

func HasItem(perso *Personnage, item string) bool {
	for _, i := range perso.Inventaire.Items {
		if i == item {
			return true
		}
	}
	return false
}

func CountItem(perso *Personnage, item string) int {
	count := 0
	for _, i := range perso.Inventaire.Items {
		if i == item {
			count++
		}
	}
	return count
}
