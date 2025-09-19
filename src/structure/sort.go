package RED

import (
	"fmt"
)

func InfoSort(perso *Personnage) {
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║           📚 SORT DU JOUEUR                ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	for _, item := range perso.Skill {
		if len(item) > 36 {
			item = item[:33] + "..."
		}
		fmt.Printf("║ • %-38s   ║\n", item)
	}

	fmt.Println("║ 🔙 Retour au Menu Principal [X]            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func SpellBookFeu(perso *Personnage) {
	aLeLivre := false
	for _, item := range perso.Inventaire.Items {
		if item == "Livre de Sort : Boule de Feu" {
			aLeLivre = true
			break
		}
	}

	if !aLeLivre {
		fmt.Println("❌ Vous n'avez pas de 'Livre de Sort : Boule de Feu' dans votre inventaire.")
		return
	}

	sort := "Boule de feu"
	for _, s := range perso.Skill {
		if s == sort {
			fmt.Println("❌ Ce sort est déjà appris.")
			return
		}
	}

	perso.Skill = append(perso.Skill, sort)
	fmt.Println("🔥 Vous avez appris le sort :", sort)

	nouvelInventaire := []string{}
	for _, item := range perso.Inventaire.Items {
		if item != "Livre de Sort : Boule de Feu" {
			nouvelInventaire = append(nouvelInventaire, item)
		}
	}
	perso.Inventaire.Items = nouvelInventaire
	fmt.Println("📖 Le 'Livre de Sort : Boule de Feu' a été consommé.")
}
func SpellBookInv(perso *Personnage) {
	aLeLivre := false
	for _, item := range perso.Inventaire.Items {
		if item == "Livre de Sort : UP Inventaire" {
			aLeLivre = true
			break
		}
	}
	if !aLeLivre {
		fmt.Println("❌ Vous n'avez pas de 'Livre de Sort : UP Inventaire' dans votre inventaire.")
		return
	}

	nouvelInventaire := []string{}
	for _, item := range perso.Inventaire.Items {
		if item != "Livre de Sort : UP Inventaire" {
			nouvelInventaire = append(nouvelInventaire, item)
		}
	}
	perso.Inventaire.Items = nouvelInventaire
	fmt.Println("📖 Le 'Livre de Sort : UP Inventaire' a été consommé.")
}
