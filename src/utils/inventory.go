package utils

import (
	"fmt"
	"strconv"
	"time"
)

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

func (p *Personnage) FindObject(nameObject string) bool {
	for _, valeur := range p.inventory {
		if valeur.nomItem == nameObject {
			return true
		}
	}
	return false
}

func (p *Personnage) AddInventory(nom string, quantite int, desc string) {
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

func (p *Personnage) LimitInventory() bool {
	return len(p.inventory) < 10
}

func (p *Personnage) TakePot() {
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

func (m *Monstre) PoisonPot() {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		if m.points_vie_actuels > 0 {
			m.points_vie_actuels -= 5
			if m.points_vie_actuels < 0 {
				break
			}
		}
		fmt.Println("[-5 POISON] Points de vie de l'ennemi : ", m.points_vie_actuels, "/", m.points_vie_maximum)
	}
}

func (m *Monstre) BouleDeFeu() {
	m.points_vie_actuels -= 20
	fmt.Println("Vie restante ennemi", m.points_vie_actuels, "pv")
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
			p.TakePot()
			break

		case "Chapeau Magique":
			if p.FindObject("Chapeau Magique") {
				p.RemoveObject("Chapeau Magique")             // ICIIIIIIII
				p.Equipement.tete.nomItem = "Chapeau Magique" // Add stuff
				p.pointVieMax += 10
				p.pointVieActual += 10
				ClearConsole()
				fmt.Println("Vous avez équipé le chapeau magique")
				p.Inventory()

			}

		case "Tunique Volante":
			if p.FindObject("Tunique Volante") {
				p.RemoveObject("Tunique Volante")
				p.Equipement.corps.nomItem = "Tunique Volante" // Add stuff
				p.pointVieMax += 25
				p.pointVieActual += 25
				ClearConsole()
				fmt.Println("Vous avez équipé la tunique Volante")
				p.Inventory()
			}

		case "Bottes de Sorcier":
			if p.FindObject("Bottes de Sorcier") {
				p.RemoveObject("Bottes de Sorcier")
				p.Equipement.pied.nomItem = "Bottes de Sorcier" // Add stuff
				p.pointVieMax += 15
				p.pointVieActual += 15
				ClearConsole()
				fmt.Println("Vous avez equipé les Bottes de sorcier")
				p.Inventory()
			}

		case "Retour ":
			p.Menu()

		}

	}
}
