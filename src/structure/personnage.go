package RED

import (
	RED "RED/menu"
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
	Inventaire []string
}

func InitCharacter(nom string, classe string, niveau int, pvMax int, pvActuels int, inventaire []string, skill []string, gold int) Personnage {
	return Personnage{
		Nom:        nom,
		Classe:     classe,
		Niveau:     niveau,
		PvMax:      pvMax,
		PvActuels:  pvActuels,
		Gold:       gold,
		Skill:      skill,
		Inventaire: inventaire,
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
		fmt.Println("║ 1. ⚔️   Humains             ║")
		fmt.Println("║ 2. 🧙  Elfes               ║")
		fmt.Println("║ 3. 🏹  Nains               ║")
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
	}
	if classe == classes[1] {
		pvMax = 80
	}
	if classe == classes[2] {
		pvMax = 120
	}
	gold := 100
	pvActuels := pvMax / 2
	inventaire := []string{}
	skill := []string{"Coup de poing"}

	personnage := InitCharacter(nom, classe, niveau, pvMax, pvActuels, inventaire, skill, gold)

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

func SpellBook(perso *Personnage) {
	// Vérifie si le joueur a un "Livre de Sort : Boule de Feu"
	aLeLivre := false
	for _, item := range perso.Inventaire {
		if item == "Livre de Sort : Boule de Feu" {
			aLeLivre = true
			break
		}
	}

	if !aLeLivre {
		fmt.Println("❌ Vous n'avez pas de 'Livre de Sort : Boule de Feu' dans votre inventaire.")
		return
	}

	// Vérifie si le sort est déjà appris
	sort := "Boule de feu"
	for _, s := range perso.Skill {
		if s == sort {
			fmt.Println("❌ Ce sort est déjà appris.")
			return
		}
	}

	// Apprendre le sort
	perso.Skill = append(perso.Skill, sort)
	fmt.Println("🔥 Vous avez appris le sort :", sort)

	// Retirer le livre de l'inventaire après utilisation
	nouvelInventaire := []string{}
	for _, item := range perso.Inventaire {
		if item != "Livre de Sort : Boule de Feu" {
			nouvelInventaire = append(nouvelInventaire, item)
		}
	}
	perso.Inventaire = nouvelInventaire
	fmt.Println("📖 Le 'Livre de Sort : Boule de Feu' a été consommé.")
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
	fmt.Printf("║ 💰 Gold : %-26d       ║\n", perso.Gold)
	fmt.Println("║ 🔙 Retour au Menu Principal (1)            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

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

	fmt.Println("║ 🔙 Retour au Menu Principal (1)            ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Println("👉 Votre choix :")
}

func IsDead(perso *Personnage) {
	if perso.PvActuels <= 0 {
		fmt.Println("💀", perso.Nom, "est mort...")
		perso.PvActuels = perso.PvMax / 2
		fmt.Println("✨", perso.Nom, "a été ressuscité avec", perso.PvActuels, "PV.")
	}
}
