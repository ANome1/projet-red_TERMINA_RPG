package RED

import (
	REDM "RED/menu"
	"fmt"
)

type Equipement struct {
	Nom       string
	Categorie string
	Effet     string
	ValEffet  int
}

type SlotsEquipement struct {
	Tete  *Equipement
	Torse *Equipement
	Pieds *Equipement
}

var Equipements = []Equipement{
	{"Chapeau de l'aventurier", "Tête", "PV Max", 10},
	{"Tunique de l'aventurier", "Torse", "PV Max", 25},
	{"Bottes de l'aventurier", "Pieds", "PV Max", 15},
}

func Equiper(perso *Personnage, item Equipement) {
	switch item.Categorie {
	case "Tête":
		if perso.Equipement.Tete != nil {
			fmt.Println("❌ Vous avez déjà un équipement pour cette catégorie.")
			return
		}
		perso.Equipement.Tete = &item
	case "Torse":
		if perso.Equipement.Torse != nil {
			fmt.Println("❌ Vous avez déjà un équipement pour cette catégorie.")
			return
		}
		perso.Equipement.Torse = &item
	case "Pieds":
		if perso.Equipement.Pieds != nil {
			fmt.Println("❌ Vous avez déjà un équipement pour cette catégorie.")
			return
		}
		perso.Equipement.Pieds = &item
	}

	if item.Effet == "PV Max" {
		perso.PvMax += item.ValEffet
		perso.PvActuels += item.ValEffet
	}
	fmt.Println("✅ Équipement équipé :", item.Nom)
}

func Desequiper(perso *Personnage, categorie string) {
	var item *Equipement
	switch categorie {
	case "Tête":
		item = perso.Equipement.Tete
		perso.Equipement.Tete = nil
	case "Torse":
		item = perso.Equipement.Torse
		perso.Equipement.Torse = nil
	case "Pieds":
		item = perso.Equipement.Pieds
		perso.Equipement.Pieds = nil
	}

	if item == nil {
		fmt.Println("❌ Aucun équipement à déséquiper pour cette catégorie.")
		return
	}

	if item.Effet == "PV Max" {
		perso.PvMax -= item.ValEffet
		if perso.PvActuels > perso.PvMax {
			perso.PvActuels = perso.PvMax
		}
	}
	fmt.Println("✅ Vous avez retiré :", item.Nom)
}

func PoisonPot(perso *Personnage) {
	if !HasItem(perso, "Potion de poison") {
		fmt.Println("❌ Vous n'avez pas de Potion de Poison dans votre inventaire.")
		return
	}
	RemoveInventory(perso, "Potion de poison")

	fmt.Println("\n╔═════════════════════════════════════════════════╗")
	fmt.Println("║ ☠️  Vous avez utilisé une Potion de Poison !     ║")
	for i := 1; i <= 3; i++ {
		REDM.Pause(1)
		perso.PvActuels -= 10
		if perso.PvActuels < 0 {
			perso.PvActuels = 0
		}
		fmt.Printf("║ 💀 Dégâts empoisonnés :%-25d║\n", 10*i)
		fmt.Printf("║ ❤️  PV actuels :%-28d     ║\n", perso.PvActuels)
	}
	fmt.Println("╚═════════════════════════════════════════════════╝")
	if perso.PvActuels == 0 {
		fmt.Println("Ne cherzhez pas la facité en mettant fin à vos jour !")
		IsDead(perso)
	}
}

func TakePot(perso *Personnage) {
	for i, v := range perso.Inventaire.Items {
		if v == "Potion de soin" {
			if perso.PvActuels < perso.PvMax {
				perso.PvActuels += 50
				if perso.PvActuels > perso.PvMax {
					perso.PvActuels = perso.PvMax
				}
				Inv := []string{}
				for j, item := range perso.Inventaire.Items {
					if j != i {
						Inv = append(Inv, item)
					}
				}
				perso.Inventaire.Items = Inv

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

func Forger(perso *Personnage, equip Equipement, materiaux map[string]int) {
	for item, qte := range materiaux {
		if CountItem(perso, item) < qte {
			fmt.Println("❌ Vous n'avez pas assez de", item)
			return
		}
	}

	if perso.Gold < 5 {
		fmt.Println("❌ Vous n'avez pas assez d'or pour forger cet équipement (5 Po requis)")
		return
	}

	for item, qte := range materiaux {
		for i := 0; i < qte; i++ {
			RemoveInventory(perso, item)
		}
	}

	perso.Gold -= 5
	REDM.ClearTerminal()
	fmt.Println("\n╔══════════════════════════════════════════════════╗")
	fmt.Println("║ Le forgeron travaille, veuillez patienter...  ⚒️  ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
	REDM.Pause(5)
	REDM.ClearTerminal()

	AddInventory(perso, equip.Nom)
	fmt.Println("✅ Le forgeron a forgé :", equip.Nom, "→ ajouté à votre inventaire")
	fmt.Println("💰 5 Po ont été dépensés pour le crafting")
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
