package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func ClearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

type Sorts struct {
	nomSort     string
	description string
	damage      int
	quantite    int
}

type Objet struct {
	nomItem     string
	description string
	quantite    int
}

type Personnage struct {
	nom            string
	class          string
	level          int
	pointVieMax    int
	pointVieActual int
	inventory      []Objet
	coins          int
	skill          []Sorts
	Equipement     Equipement
}

type Equipement struct {
	tete  Objet
	corps Objet
	pied  Objet
}

type Monstre struct {
	nom                string
	points_vie_maximum int
	points_vie_actuels int
	pointsAttaque      int
}

var m = Monstre{"Gobelin", 40, 40, 20}

func (m *Monstre) InitGoblin(nom string, points_vie_maximum int, points_vie_actuels int, pointsAttaque int) {
	m.nom = nom
	m.points_vie_maximum = points_vie_maximum
	m.points_vie_actuels = points_vie_actuels
	m.pointsAttaque = pointsAttaque
}

func (p *Personnage) Init(nom string, class string, level int, pointVieMax int, pointVieActual int, coins int, skill []Sorts) {
	p.nom = nom
	p.class = class
	p.level = level
	p.pointVieMax = pointVieMax
	p.pointVieActual = pointVieActual
	p.coins = coins
	p.skill = skill
}

func (o *Objet) Init1(nomItem string, description string, quantite int) {
	o.nomItem = nomItem
	o.description = description
	o.quantite = quantite
}

func (s *Sorts) Init2(nomSort string, description string, damage int, quantite int) {
	s.nomSort = nomSort
	s.description = description
	s.damage = damage
	s.quantite = quantite
}

func (p *Personnage) FindObject(nameObject string) bool {
	for _, valeur := range p.inventory {
		if valeur.nomItem == nameObject {
			return true
		}
	}
	return false
}

func (p *Personnage) RemoveObject(NomItem string) {
	for index, valeur := range p.inventory {
		if valeur.nomItem == NomItem {
			if valeur.quantite == 1 {
				p.inventory = append(p.inventory[:index], p.inventory[index+1:]...)
			} else {
				valeur.quantite--
			}
		}
	}
}

func (p *Personnage) limitInventory() bool {
	return len(p.inventory) < 10
}

func (p *Personnage) Inventory() {
	for i := 0; i < len(p.inventory); i++ {
		fmt.Println("["+strconv.Itoa(i+1)+"] ", p.inventory[i].nomItem, " : ", p.inventory[i].quantite)

	}
	fmt.Println("[" + strconv.Itoa(len(p.inventory)+1) + "] Retour")
	choix, _ := Inputint()
	if choix == len(p.inventory)+1 {
		ClearConsole()
		p.Menu()
		// retour menu
	} else {
		var objet Objet = p.inventory[choix-1]
		switch objet.nomItem {
		case "Potions de vie":
			p.takePot()
			break

		case "Chapeau Magique":
			if p.FindObject("Chapeau Magique") {
				p.RemoveObject("Chapeau Magique")             // ICIIIIIIII
				p.Equipement.tete.nomItem = "Chapeau Magique" // Add stuff
				p.pointVieMax += 10
				fmt.Println("Vous avez équipé le chapeau magique")
				p.Inventory()

			}
			break

		case "Tunique Volante":
			if p.FindObject("Tunique Volante") {
				p.RemoveObject("Tunique Volante")
				p.Equipement.corps.nomItem = "Tunique Volante" // Add stuff
				p.pointVieMax += 25
				fmt.Println("Vous avez équipé la tunique Volante")
				p.Inventory()
			}
			break

		case "Bottes de sorcier":
			if p.FindObject("Bottes de Sorcier") {
				p.RemoveObject("Bottes de Sorcier")
				p.Equipement.pied.nomItem = "Bottes de Sorcier" // Add stuff
				p.pointVieMax += 15
				fmt.Println("Vous avez equipé les Bottes de sorcier")
				p.Inventory()
			}
			break

		case "Retour ":
			p.Menu()

		}

	}
}

func (p *Personnage) Skill() {
	fmt.Println(p.skill)
}

func (p *Personnage) deadPersonnage() {
	if p.pointVieActual < 0 {
		fmt.Println("Vous etes mort !!!")
		p.pointVieActual = p.pointVieActual / 2
		fmt.Println("Vous etes résuscité avec la moitié de votre vie")
		p.Menu()
	}
}
func (m *Monstre) deadMonstre() {
	if m.points_vie_actuels < 0 {
		fmt.Println("Le monstre est mort !!!")
	}
}

func (p *Personnage) takePot() {
	if p.inventory[0].quantite > 0 {
		p.inventory[0].quantite -= 1
		if p.pointVieActual < 50 {
			p.pointVieActual += 50
		} else if p.pointVieActual >= 50 && p.pointVieActual < 100 {
			p.pointVieActual = p.pointVieMax
			fmt.Println("Vous avez utilisé la potion de vie")
		} else {
			fmt.Println("Impossible de boire la potion !")
		}
	} else {
		fmt.Println(" Vous n'avez pas de potion !")
	}

}

func (m *Monstre) poisonPot() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		if m.points_vie_actuels > 0 {
			m.points_vie_actuels -= 5
			if m.points_vie_actuels < 0 {
				break
			}
		}
		fmt.Println("[-10 POISON] Points de vie de l'ennemi : ", m.points_vie_actuels, "/", m.points_vie_maximum)
	}
}

func (p *Personnage) enoughMoney(prix int) bool {
	if p.coins-prix < 0 {
		return false
	} else {
		return true
	}
}

func (m *Monstre) PlayerRound(p *Personnage) {
	fmt.Println("1. Poison")
	attaque, _ := Inputint()
	switch attaque {
	case 1:
		fmt.Println("Vous avez utiliser le poison sur le monstre")
		m.poisonPot()
	}

}

func (p *Personnage) enemieRound() {
	fmt.Println("Griffure")
	p.pointVieActual -= m.pointsAttaque
	fmt.Println("Vie restante", p.pointVieActual)
}

func (m *Monstre) trainFight(p *Personnage) {
	for i := 0; p.pointVieActual > 0 || m.points_vie_actuels > 0; i++ {
		m.PlayerRound(p)
		m.deadMonstre()
		p.enemieRound()
		p.deadPersonnage()
	}

}

func (p *Personnage) marchand() {
	fmt.Println("Bienvenue chez le marchand que voulez vous acheter ? \n")
	fmt.Println("1. Potion vie : 3 PO")
	fmt.Println("2. Potion poison : 6 PO")
	fmt.Println("3. Livre de Sort : Boule de feu : 25 PO")
	fmt.Println("4. Chapeau Magique : 5 PO")
	fmt.Println("5. Tunique Volante : 8 PO")
	fmt.Println("6. Bottes de sorcier : 5 PO")
	fmt.Println("7. Retour")
	fmt.Printf("------------------ \n")
	fmt.Printf("------------------ \n")
	fmt.Printf("Votre solde actuel : %d\n", p.coins)

	choix, _ := Inputint()

	switch choix {
	case 1:
		nom := "Potions de vie"
		desc := "Donne 50 PV"
		quantite := 1
		prix := 3

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté une potion de vie ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire !")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}
		break
	case 2:
		fmt.Println("Vous avez acheté une potion de poison ")
		nom := "Potion de poison"
		desc := "Enleve 20% des PV"
		quantite := 1
		prix := 6

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté une potion de poison ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}
		break

	case 3:
		fmt.Println("Vous avez acheté une boule de feu ")
		nom := "Boule de feu"
		desc := "Boule de feu qui enleve 25pv"
		quantite := 1
		prix := 25

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté un livre de sort ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}
		break

	case 4:
		nom := "Chapeau Magique"
		desc := ""
		quantite := 1
		prix := 5 // Add PV

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté le chapeau magique ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}
		break
	case 5:
		nom := "Tunique Volante"
		desc := ""
		quantite := 1
		prix := 8

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté la tunique volante ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}
		break // Add PV

	case 6:
		// Add PV
		nom := " Bottes de sorcier"
		desc := ""
		quantite := 1
		prix := 5

		if p.enoughMoney(prix) {
			if p.limitInventory() {
				p.addInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				fmt.Println("Vous avez acheté les bottes de l'aventurier ")
				p.marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.marchand()
			}
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()

		}
		break
	case 7:
		ClearConsole()
		p.Menu()
		break
	}
}

func main() {
	var p1 Personnage
	p1.Init("Harry", "Sorcier", 5, 100, 100, 100, []Sorts{{"Coup de poing", "Auto attaque de base", 10, 1}})
	ClearConsole()
	p1.Menu()
}

func (p *Personnage) addInventory(nom string, quantite int, desc string) {
	existe := false
	var index int
	for i, item := range p.inventory {
		if item.nomItem == nom {
			existe = true
			index = i
		}
	}
	if existe {
		p.inventory[index].quantite += quantite
	} else {
		nouvelItem := Objet{nomItem: nom, description: desc, quantite: quantite}
		p.inventory = append(p.inventory, nouvelItem)
	}
}

func (p *Personnage) displayInfo() {
	fmt.Printf("Nom: %s\n", p.nom)
	fmt.Printf("Class: %s\n", p.class)
	fmt.Printf("Niveau: %d\n", p.level)
	fmt.Printf("Point_vie_max: %d\n", p.pointVieMax)
	fmt.Printf("Point_vie_actuel: %d\n", p.pointVieActual)
	fmt.Printf("Coins: %d\n", p.coins)
	fmt.Printf("skill: %s\n", p.skill)
	fmt.Printf("Equipement: %s\n", p.Equipement)
	fmt.Printf("------------------ \n")

	fmt.Println("\n")
}

func Inputint() (int, error) {
	fmt.Print(">> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	chiffre, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return chiffre, nil
}

func (p *Personnage) Menu() {
	fmt.Println("Veuillez vous diriger dans le menu :  \n")
	fmt.Printf("1.  Personnage \n")
	fmt.Printf("2.  Inventaire \n")
	fmt.Printf("3.  Marchand \n")
	fmt.Printf("4.  Combat \n")

	choice, _ := Inputint()

	switch choice {
	case 1:
		ClearConsole()
		p.displayInfo()
		p.Menu()

	case 2:
		fmt.Println("Affichage de l'inventaire :")
		ClearConsole()
		p.Inventory()

	case 3:
		ClearConsole()
		p.marchand()

	case 4:
		m.trainFight(p)

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")

	}
}
