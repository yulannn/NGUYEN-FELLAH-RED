package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func (p *Personnage) Menu() {
	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	color.Blue("       Veuillez vous diriger dans le menu :              ")
	fmt.Println("║                                                   ║")
	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	fmt.Println("║ 1.  Personnage                                    ║")
	fmt.Println("║ 2.  Inventaire                                    ║")
	fmt.Println("║ 3.  Skill                                         ║")
	fmt.Println("║ 4.  Marchand                                      ║")
	fmt.Println("║ 5.  Combat                                        ║")
	fmt.Println("║                                                   ║")
	fmt.Println("╚═══════════════════════════════════════════════════╝")

	choice, _ := Inputint()

	switch choice {
	case 1:
		ClearConsole()
		p.DisplayInfo()
		p.Menu()
		break

	case 2:
		fmt.Println("Affichage de l'inventaire :")
		ClearConsole()
		p.Inventory()
		break

	case 3:
		fmt.Println("Affichage des skills :")
		ClearConsole()
		p.Skill()
		p.Menu()
		break

	case 4:
		ClearConsole()
		p.Marchand()
		break

	case 5:
		m.TrainFight(p)
		break

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")
		p.Menu()
		break

	}
}
