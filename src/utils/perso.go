package utils

type Sorts struct {
	NomSort     string
	Description string
	Damage      int
	Quantite    int
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
	Equipement     Equipement
}

type Equipement struct {
	tete  Objet
	corps Objet
	pied  Objet
}

type Monstre struct {
	nom                string
	points_vie_maximum int
	points_vie_actuels int
	pointsAttaque      int
}

var m = Monstre{"Gobelin", 40, 40, 20}

func (m *Monstre) InitGoblin(nom string, points_vie_maximum int, points_vie_actuels int, pointsAttaque int) {
	m.nom = nom
	m.points_vie_maximum = points_vie_maximum
	m.points_vie_actuels = points_vie_actuels
	m.pointsAttaque = pointsAttaque
}

func (o *Objet) Init1(nomItem string, description string, quantite int) {
	o.nomItem = nomItem
	o.description = description
	o.quantite = quantite
}

func (s *Sorts) Init2(nomSort string, description string, damage int, quantite int) {
	s.NomSort = nomSort
	s.Description = description
	s.Damage = damage
	s.Quantite = quantite
}

func (p *Personnage) Initialize(nom string, class string, level int, pointVieMax int, pointVieActual int, coins int, skill []Sorts) {
	p.nom = nom
	p.class = class
	p.level = level
	p.pointVieMax = pointVieMax
	p.pointVieActual = pointVieActual
	p.coins = coins
	p.skill = skill
}
