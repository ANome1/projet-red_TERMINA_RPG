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
}
