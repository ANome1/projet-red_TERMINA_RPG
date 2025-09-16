package RED

import (
	"fmt"
)

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PvMax      int
	PvActuels  int
	Inventaire []string
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string) Personnage {
	return Personnage{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PvMax:      pvMax,
		PvActuels:  pvActuels,
		Inventaire: Inventaire{Items: inventaire, Max: 10}.Items,
	}
}

func DisplayInfo(perso Personnage) {
	fmt.Println("=== Informations du personnage ===")
	fmt.Println("Nom :", perso.Nom)
	fmt.Println("Classe :", perso.Classe)
	fmt.Println("Niveau :", perso.Niveau)
	fmt.Println("PV actuels :", perso.PvActuels)
	fmt.Println("PV max :", perso.PvMax)
	fmt.Println("1. Retour au Menu Principal")
}

func IsDead(perso *Personnage) {
	if perso.PvActuels <= 0 {
		fmt.Println("💀", perso.Nom, "est mort...")
		perso.PvActuels = perso.PvMax / 2
		fmt.Println("✨", perso.Nom, "a été ressuscité avec", perso.PvActuels, "PV.")
	}
}
