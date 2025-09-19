package RED

import (
	RED "RED/menu"
	REDM "RED/menu"
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Personnage struct {
	Nom        string
	Classe     string
	Niveau     int
	PvMax      int
	PvActuels  int
	Gold       int
	Skill      []string
	Inventaire Inventaire
	Equipement SlotsEquipement
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire Inventaire, skill []string, gold int, equipement SlotsEquipement) Personnage {
	return Personnage{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PvMax:      pvMax,
		PvActuels:  pvActuels,
		Gold:       gold,
		Skill:      skill,
		Inventaire: inventaire,
		Equipement: equipement,
	}
}

func CharacterCreation() Personnage {
	reader := bufio.NewReader(os.Stdin)

	var nom string
	for {
		fmt.Print("🧑 Entrez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if estNomValide(input) {
			nom = formaterNom(input)
			break
		} else {
			fmt.Println("❌ Le nom ne doit contenir que des lettres.")
		}
	}

	classes := []string{"Humains", "Elfes", "Nains"}
	var choixClasse int
	var choixValide bool

	for !choixValide {
		fmt.Println("\n╔════════════════════════════╗")
		fmt.Println("║  🔰 Choisissez une classe  ║")
		fmt.Println("╠════════════════════════════╣")
		fmt.Println("║ [1] ⚔️   Humains            ║")
		fmt.Println("║ [2] 🧙  Elfes              ║")
		fmt.Println("║ [3] 🏹  Nains              ║")
		fmt.Println("╚════════════════════════════╝")
		fmt.Print("👉 Votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1", "2", "3":
			choixClasse = int(input[0] - '1')
			choixValide = true
		default:
			fmt.Println("❌ Choix invalide. Veuillez entrer un nombre entre 1 et 3.")
		}
	}

	classe := classes[choixClasse]
	niveau := 1
	var pvMax int
	if classe == classes[0] {
		pvMax = 100
	} else if classe == classes[1] {
		pvMax = 80
	} else if classe == classes[2] {
		pvMax = 120
	}

	gold := 100
	pvActuels := pvMax / 2
	inventaire := Inventaire{
		Items:      []string{},
		Max:        10,
		CptUpgrade: 0,
	}

	skill := []string{"Coup de poing"}
	equipement := SlotsEquipement{}

	personnage := InitCharacter(nom, classe, niveau, pvMax, pvActuels, inventaire, skill, gold, equipement)

	RED.ClearTerminal()
	fmt.Println("\n╔════════════════════════════════════╗")
	fmt.Println("║   🎉 PERSONNAGE CRÉÉ AVEC SUCCÈS   ║")
	fmt.Println("╠════════════════════════════════════╣")
	fmt.Printf("║ 🔤 Nom    : %-22s ║\n", personnage.Nom)
	fmt.Printf("║ 🧙 Classe : %-22s ║\n", personnage.Classe)
	fmt.Println("╚════════════════════════════════════╝")

	fmt.Print("\n✨ Appuyez sur Entrée pour commencer votre aventure...")
	reader.ReadString('\n')

	RED.ClearTerminal()
	return personnage
}

func estNomValide(nom string) bool {
	for _, r := range nom {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func formaterNom(nom string) string {
	if len(nom) == 0 {
		return nom
	}
	nom = strings.ToLower(nom)
	return strings.ToUpper(string(nom[0])) + nom[1:]
}

func DisplayInfo(perso Personnage) {
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║        📋 INFORMATIONS DU PERSONNAGE       ║")
	fmt.Println("╠════════════════════════════════════════════╣")
	fmt.Printf("║ 🔤 Nom    : %-30s ║\n", perso.Nom)
	fmt.Printf("║ 🧙 Classe : %-30s ║\n", perso.Classe)
	fmt.Printf("║ 📈 Niveau : %-30d ║\n", perso.Niveau)
	fmt.Printf("║ ❤️  PVActuels : %-26d  ║\n", perso.PvActuels)
	fmt.Printf("║ 💖 PVMax : %-26d      ║\n", perso.PvMax)
	fmt.Printf("║ 💰 Po : %-26d         ║\n", perso.Gold)
	fmt.Println("║ 🔙 Retour au Menu Principal [X]            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func IsDead(perso *Personnage) bool {
	if perso.PvActuels <= 0 {
		fmt.Println("\n💀 Vous êtes mort, vous allez réapparaitre dans 5 secondes...")
		REDM.Pause(5)
		perso.PvActuels = perso.PvMax / 2
		fmt.Printf("✨ Vous ressuscitez avec %d PV !\n", perso.PvActuels)
		REDM.Pause(2)
		return true
	}
	return false
}
