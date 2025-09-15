package RED

type Items struct {
	Items  []string
	Max    int
	Effet  []string
	Valeur int
}

func TakePotion(perso *Personnage, item string) {
	for i, it := range perso.Inventaire {
		if it == item {
			if item == "Potion de soin" {
				perso.PvActuels += 50
			}
			if perso.PvActuels > perso.PvMax {
				perso.PvActuels = perso.PvMax
			}
			// Retirer l'objet de l'inventaire
			perso.Inventaire = append(perso.Inventaire[:i], perso.Inventaire[i+1:]...)
			break
		}
	}
}
