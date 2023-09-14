package piscine

import "fmt"

type Personnage struct {
	Nom              string
	Classe           string
	Niveau           int
	Point_vie_max    int
	Point_vie_actuel int
	Inventaire       string
}

func Init(nom string, classe string, niveau int, Point_vie_max int, Point_vie_actuel int, inventaire string) Personnage {
	personnage := Personnage{
		Nom:              nom,
		Classe:           classe,
		Niveau:           niveau,
		Point_vie_max:    Point_vie_max,
		Point_vie_actuel: Point_vie_actuel,
		Inventaire:       inventaire,
	}
	return personnage
}

func main() {
	NewPerso := Init("Harry", "Sorcier", 5, 100, 80, "Pleins")
	fmt.Printf("Nom: %s\n", NewPerso.Nom)
	fmt.Printf("Class: %s\n", NewPerso.Classe)
	fmt.Printf("Niveau: %d\n", NewPerso.Niveau)
	fmt.Printf("Point_vie_max: %d\n", NewPerso.Point_vie_max)
	fmt.Printf("Point_vie_actuel: %d\n", NewPerso.Point_vie_actuel)
	fmt.Printf("Inventaire: %s\n", NewPerso.Inventaire)
}
