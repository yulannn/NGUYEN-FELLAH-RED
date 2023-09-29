package utils

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func (m *Monstre) PlayerRound(p *Personnage) {
	fmt.Println("              _,._      ")
	fmt.Println("  .||,       /_ _\\     ")
	fmt.Println(" \\.`',/      |'L'| |    ")
	fmt.Println(" = ,. =      | -,| L    ")
	fmt.Println(" / || \\    ,-'\"/,'`.   ")
	fmt.Println("   ||     ,'   `,,. `.  ")
	fmt.Println("   ,|____,' , ,;' \\| |  ")
	fmt.Println("  (3|\\    _/|/'   _| |  ")
	fmt.Println("   ||/,-''  | >-'' _,\\ ")
	fmt.Println("   ||'      ==\\ ,-'  ,' ")

	color.Blue("Quelle attaque voulez vous utiliser ")
	index := 0
	for _, i := range p.skill {
		fmt.Printf("%d. %s \n", index+1, i.NomSort)
		index++
	}
	attaque, _ := Inputint()
	if attaque <= len(p.skill) {
		skill := p.skill[attaque-1]
		switch skill.NomSort {
		case "Coup de poing":
			m.points_vie_actuels -= skill.Damage
			ClearConsole()
			color.Green("Vous avez infligé un coup de poing ")
			if m.points_vie_actuels < 0 {
				m.points_vie_actuels = 0
			}
			fmt.Println("Vie restante ennemi", m.points_vie_actuels, "pv")
			fmt.Println("------------------------------")
			break

		case "Potion de poison":
			color.Green("Vous avez utilisé la potion de poison ")
			fmt.Println("Mana restant :", p.mana)
			m.PoisonPot()
			break

		case "Boule de feu":
			color.Green("Vous avez utilisé la boule de feu ")
			fmt.Println("Mana restant :", p.mana)
			m.BouleDeFeu()

		}

	}

}

func (p *Personnage) EnemieRound() {
	fmt.Println("       .-.       ")
	fmt.Println("      ( \")      ")
	fmt.Println("   /\\_.' '._/\\   ")
	fmt.Println("   |         |   ")
	fmt.Println("    \\       /    ")
	fmt.Println("     \\    /`     ")
	fmt.Println("    (__)  /      ")
	fmt.Println("   `.__.'        ")
	color.Red("Le détraqueur vous a infligé Griffure")
	p.pointVieActual -= m.pointsAttaque
	fmt.Println("Il vous reste ", p.pointVieActual, "pv")
	fmt.Println("------------------------------")
	time.Sleep(2000 * time.Millisecond)
}

func (m *Monstre) TrainFight(p *Personnage) {
	for i := 0; p.pointVieActual > 0 || m.points_vie_actuels > 0; i++ {
		if p.initiative > m.initiative {
			m.PlayerRound(p)
			m.DeadMonstre(p)
			time.Sleep(2000 * time.Millisecond)
			p.EnemieRound()
			p.DeadPersonnage()
		} else {
			p.EnemieRound()
			p.DeadPersonnage()
			time.Sleep(2000 * time.Millisecond)
			m.PlayerRound(p)
			m.DeadMonstre(p)
		}

	}

}

func (p *Personnage) DeadPersonnage() {
	if p.pointVieActual <= 0 {
		fmt.Println("Vous etes mort !!!")
		p.pointVieActual = p.pointVieActual / 2
		fmt.Println("Vous etes résuscité avec la moitié de votre vie")
		time.Sleep(3000 * time.Millisecond)
		ClearConsole()
		p.Menu()
	}
}

func (m *Monstre) DeadMonstre(p *Personnage) {
	if m.points_vie_actuels <= 0 {
		fmt.Println("Le détraqueur est mort !!!")
		p.xp += 100
		p.Niveau()
		fmt.Println("Vous avez gagné 100xp vous passez level 2 ")
		time.Sleep(3000 * time.Millisecond)
		ClearConsole()
		p.Menu()
	}
}
