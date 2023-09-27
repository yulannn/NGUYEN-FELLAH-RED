package utils

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func (m *Monstre) PlayerRound(p *Personnage) {
	color.Blue("Quelle attaque voulez vous utiliser ")
	index := 0
	for _, i := range p.skill {
		fmt.Printf("%d. %s \n", index+1, i.NomSort)
		index++
	}
	for _, i := range p.inventory {
		fmt.Printf("%d. %s \n", index+1, i.nomItem)
		index++

	}
	attaque, _ := Inputint()
	if attaque <= len(p.skill) {
		skill := p.skill[attaque-1]
		switch skill.NomSort {
		case "Coup de poings":
			m.points_vie_actuels -= skill.Damage
			color.Green("Vous avez infligé un coup de poing ")
			if m.points_vie_actuels < 0 {
				m.points_vie_actuels = 0
			}
			fmt.Println("Vie restante ennemi", m.points_vie_actuels, "pv")
			fmt.Println("------------------------------")
			break
		}
	} else {
		sort := attaque - len(p.skill)
		objet := p.inventory[sort-1]
		switch objet.nomItem {
		case "Potion de poison":
			color.Green("Vous avez utilisé la potion de poison ")
			m.PoisonPot()
			break
		case "Boule de feu":
			color.Green("Vous avez utilisé la boule de feu ")
			m.BouleDeFeu()

		}

	}

}

func (p *Personnage) EnemieRound() {
	color.Red("Le monstre vous a infligé Griffure")
	p.pointVieActual -= m.pointsAttaque
	fmt.Println("Il vous reste ", p.pointVieActual, "pv")
	fmt.Println("------------------------------")
	time.Sleep(1000 * time.Millisecond)
}

func (m *Monstre) TrainFight(p *Personnage) {
	for i := 0; p.pointVieActual > 0 || m.points_vie_actuels > 0; i++ {
		m.PlayerRound(p)
		m.DeadMonstre(p)
		time.Sleep(1000 * time.Millisecond)
		p.EnemieRound()
		p.DeadPersonnage()
	}

}

func (p *Personnage) DeadPersonnage() {
	if p.pointVieActual <= 0 {
		fmt.Println("Vous etes mort !!!")
		p.pointVieActual = p.pointVieActual / 2
		fmt.Println("Vous etes résuscité avec la moitié de votre vie")
		time.Sleep(2500 * time.Millisecond)
		ClearConsole()
		p.Menu()
	}
}

func (m *Monstre) DeadMonstre(p *Personnage) {
	if m.points_vie_actuels <= 0 {
		fmt.Println("Le monstre est mort !!!")
		time.Sleep(2500 * time.Millisecond)
		ClearConsole()
		p.Menu()
	}
}
