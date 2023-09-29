package utils

type Sorts struct {
	NomSort  string
	Damage   int
	Quantite int
	ManaCost int
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
	initiative     int
	xp             int
	xp_max         int
	mana           int
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
	initiative         int
}

var m = Monstre{"Fantome", 40, 40, 10, 5}

func (m *Monstre) InitMonstre(nom string, points_vie_maximum int, points_vie_actuels int, pointsAttaque int, initiative int) {
	m.nom = nom
	m.points_vie_maximum = points_vie_maximum
	m.points_vie_actuels = points_vie_actuels
	m.pointsAttaque = pointsAttaque
	m.initiative = initiative
}

func (o *Objet) Init1(nomItem string, description string, quantite int) {
	o.nomItem = nomItem
	o.description = description
	o.quantite = quantite
}

func (s *Sorts) Init2(nomSort string, damage int, quantite int, ManaCost int) {
	s.NomSort = nomSort
	s.Damage = damage
	s.Quantite = quantite
	s.ManaCost = ManaCost
}

func (p *Personnage) Initialize(nom string, class string, level int, pointVieMax int, pointVieActual int, coins int, skill []Sorts, initiative int, xp int, xp_max int, mana int) {
	p.nom = nom
	p.class = class
	p.level = level
	p.pointVieMax = pointVieMax
	p.pointVieActual = pointVieActual
	p.coins = coins
	p.skill = skill
	p.initiative = initiative
	p.xp = xp
	p.xp_max = xp_max
	p.mana = mana
}

func (p *Personnage) Niveau() {
	if p.xp == p.xp_max {
		p.level += 1
		p.pointVieActual += 20
		p.pointVieMax += 20
		p.coins += 10
	}
}
