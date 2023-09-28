package main

import "src/utils"

func main() {
	var p1 utils.Personnage
	var sorts = []utils.Sorts{
		{
			NomSort:  "Coup de poing",
			Damage:   10,
			Quantite: 1,
		},
	}
	p1.Initialize("Son Goku", "Sorcier", 1, 100, 100, 100, sorts, 6, 0, 100, 100)
	utils.ClearConsole()
	p1.Menu()
}
