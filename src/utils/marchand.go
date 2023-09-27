package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func (p *Personnage) EnoughMoney(prix int) bool {
	if p.coins-prix < 0 {
		return false
	} else {
		return true
	}
}

func (p *Personnage) Marchand() {
	fmt.Println("╔══════════════════════════════════════════╗")
	fmt.Println("║      Bienvenue chez le marchand          ║")
	fmt.Println("╠══════════════════════════════════════════╣")
	fmt.Println("║  Que voulez-vous acheter ?               ║")
	fmt.Println("║                                          ║")
	fmt.Println("║  1. Potion de vie :      3 PO            ║")
	fmt.Println("║  2. Potion de poison :   6 PO            ║")
	fmt.Println("║  3. Livre de Sort - Boule de feu : 25 PO ║")
	fmt.Println("║  4. Chapeau Magique :    5 PO            ║")
	fmt.Println("║  5. Tunique Volante :    8 PO            ║")
	fmt.Println("║  6. Bottes de sorcier :  5 PO            ║")
	fmt.Println("║  7. Retour                               ║")
	fmt.Println("╚══════════════════════════════════════════╝")
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

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté une potion de vie ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire !")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()
		}
		break
	case 2:
		color.Green("Vous avez acheté une potion de poison ")
		nom := "Potion de poison"
		desc := "Enleve 20% des PV"
		quantite := 1
		prix := 6

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté une potion de poison ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()
		}
		break

	case 3:
		color.Green("Vous avez acheté une boule de feu ")
		nom := "Boule de feu"
		desc := "Boule de feu qui enleve 20pv"
		quantite := 1
		prix := 25

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté un livre de sort ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()
		}
		break

	case 4:
		nom := "Chapeau Magique"
		desc := ""
		quantite := 1
		prix := 5 // Add PV

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté le chapeau magique ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()
		}
		break
	case 5:
		nom := "Tunique Volante"
		desc := ""
		quantite := 1
		prix := 8

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté la tunique volante ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()
		}
		break // Add PV

	case 6:
		// Add PV
		nom := " Bottes de sorcier"
		desc := ""
		quantite := 1
		prix := 5

		if p.EnoughMoney(prix) {
			if p.LimitInventory() {
				p.AddInventory(nom, quantite, desc)
				p.coins -= prix
				ClearConsole()
				color.Green("Vous avez acheté les bottes de l'aventurier ")
				p.Marchand()
			} else {
				ClearConsole()
				fmt.Println("Vous n'avez pas assez de place dans l'inventaire ")
				p.Marchand()
			}
		} else {
			ClearConsole()
			color.Red("Vous n'avez pas assez d'argent ")
			p.Marchand()

		}
		break
	case 7:
		ClearConsole()
		p.Menu()
		break
	}
}
