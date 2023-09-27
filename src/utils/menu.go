package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func (p *Personnage) Menu() {
	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	color.Blue("     Veuillez vous diriger dans le menu :              ")
	fmt.Println("║                                                   ║")
	fmt.Println("║═══════════════════════════════════════════════════║")
	fmt.Println("║                                                   ║")
	fmt.Println("║ 1.  Personnage                                    ║")
	fmt.Println("║ 2.  Inventaire                                    ║")
	fmt.Println("║ 3.  Marchand                                      ║")
	fmt.Println("║ 4.  Combat                                        ║")
	fmt.Println("║                                                   ║")
	fmt.Println("╚═══════════════════════════════════════════════════╝")

	choice, _ := Inputint()

	switch choice {
	case 1:
		ClearConsole()
		p.DisplayInfo()
		p.Menu()

	case 2:
		fmt.Println("Affichage de l'inventaire :")
		ClearConsole()
		p.Inventory()

	case 3:
		ClearConsole()
		p.Marchand()

	case 4:
		m.TrainFight(p)

	default:
		fmt.Println("Choix invalide. Veuillez sélectionner une autre option.")

	}
}
