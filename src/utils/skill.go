package utils

import (
	"fmt"
	"strconv"
)

func (p *Personnage) AddSkill(NomSort string, Quantite int, Damage int, ManaCost int) {
	existe := false
	var index int
	for i, item := range p.skill {
		if item.NomSort == NomSort {
			existe = true
			index = i
		}
	}
	if existe {
		p.skill[index].Quantite += Quantite
	} else {
		nouvelItem := Sorts{NomSort: NomSort, Quantite: Quantite, Damage: Damage, ManaCost: ManaCost}
		p.skill = append(p.skill, nouvelItem)
	}
}

func (p *Personnage) RemoveSkill(NomSort string) {
	for index, valeur := range p.skill {
		if valeur.NomSort == NomSort {
			if valeur.Quantite == 1 {
				p.skill = append(p.skill[:index], p.skill[index+1:]...)
			} else {
				valeur.Quantite--
			}
		}
	}
}

func (p *Personnage) FindSkill(nameSkill string) bool {
	for _, valeur := range p.inventory {
		if valeur.nomItem == nameSkill {
			return true
		}
	}
	return false
}

func (p *Personnage) LimitSkill() bool {
	return len(p.skill) < 10
}

func (p *Personnage) Skill() {
	for i := 0; i < len(p.skill); i++ {
		fmt.Println("["+strconv.Itoa(i+1)+"] ", p.skill[i].NomSort, " : ", p.skill[i].Quantite)
	}
}
