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
}

func (p *Personnage) Init(nom string, class string, level int, pointVieMax int, pointVieActual int, coins int, inventory []Objet, skill []Sorts) {
	p.nom = nom
	p.class = class
	p.level = level
	p.pointVieMax = pointVieMax
	p.pointVieActual = pointVieActual
	p.inventory = inventory
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

func (p *Personnage) Inventory() {
	for i := 0; i < len(p.inventory); i++ {
		fmt.Println("["+strconv.Itoa(i+1)+"] ", p.inventory[i].nomItem, " : ", p.inventory[i].quantite)
	}
	choix4, _ := Inputint()

	switch choix4 {
	case 1:
		p.takePot()
		fmt.Println("Inventaire restant : ")
		p.Inventory()
	}

}

func (p *Personnage) Skill() {
	fmt.Println(p.skill)
}

func (p *Personnage) dead() {
	if p.pointVieActual < 0 {
		fmt.Println("Vous etes mort !!!")
	}
}

func (p *Personnage) spellBook() {
}

func (p *Personnage) takePot() {
	fmt.Println("Vous avez utilisé la potion de vie")
	if p.inventory[0].quantite > 0 {
		p.inventory[0].quantite -= 1
		if p.pointVieActual < 50 {
			p.pointVieActual += 50
		} else if p.pointVieActual >= 50 && p.pointVieActual < 100 {
			p.pointVieActual = p.pointVieMax
		} else {
			fmt.Println("Impossible de boire la potion !")
		}
	} else {
		fmt.Println(" Vous n'avez pas de potion !")
	}

}

func (p *Personnage) poisonPot() {
	vie := p.pointVieActual / 5 * 4
	limitPv := p.pointVieActual - vie
	for {
		time.Sleep(2 * time.Second)
		p.pointVieActual -= 5
		if p.pointVieActual <= limitPv {
			fmt.Println("Fin du poison")
			break
		}
	}
}

func (p *Personnage) enoughMoney(prix int) bool {
	if p.coins-prix < 0 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) marchand() {
	fmt.Println("Bienvenue chez le marchand que voulez vous acheter ? \n")
	fmt.Println("1. Potion vie : 3 PO")
	fmt.Println("2. Potion poison : 6 PO")
	fmt.Println("3. Livre de Sort : Boule de feu : 25 PO")
	fmt.Println("4. Retour")
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
			p.addInventory(nom, quantite, desc)
			p.coins -= prix
			ClearConsole()
			fmt.Println("Vous avez acheté une potion de vie ")
			p.marchand()
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}

	case 2:
		fmt.Println("Vous avez acheté une potion de poison ")
		nom := "Potion de poison"
		desc := "Enleve 20% des PV"
		quantite := 1
		prix := 6
		if p.enoughMoney(prix) {
			p.addInventory(nom, quantite, desc)
			p.coins -= prix
			ClearConsole()
			fmt.Println("Vous avez acheté une potion de poison ")
			p.marchand()
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}

	case 3:
		fmt.Println("Vous avez acheté une boule de feu ")
		nom := "Boule de feu"
		desc := "Boule de feu qui enleve 25pv"
		quantite := 1
		prix := 25
		if p.enoughMoney(prix) {
			p.addInventory(nom, quantite, desc)
			p.coins -= prix
			ClearConsole()
			fmt.Println("Vous avez acheté un livre de sort ")
			p.marchand()
		} else {
			ClearConsole()
			fmt.Println("Vous n'avez pas assez d'argent ")
			p.marchand()
		}

	case 4:
		ClearConsole()
		p.Menu()

	}
}

func main() {
	var p1 Personnage
	p1.Init("Harry", "Sorcier", 5, 100, 40, 100, []Objet{{"Potions de vie", "Donne de la vie", 0}}, []Sorts{{"Coup de poing", "Auto attaque de base", 10, 1}})
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
		p.Menu()

	case 3:
		ClearConsole()
		p.marchand()

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")

	}
}
