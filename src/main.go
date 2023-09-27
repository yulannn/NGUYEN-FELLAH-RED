package main

import "src/utils"

func main() {
	var p1 utils.Personnage
	var sorts = []utils.Sorts{
		{
			NomSort:     "Coup de poings",
			Description: "",
			Damage:      10,
			Quantite:    1,
		},
	}
	p1.Initialize("Harry", "Sorcier", 5, 100, 100, 100, sorts)
	utils.ClearConsole()
	p1.Menu()
}
