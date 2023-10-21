//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Player struct {
	name              string
	health, maxhealth uint
	energy, maxenergy uint
}

func (p Player) addHealth(health uint) Player {
	p.health += health
	if p.health > p.maxhealth {
		p.health = p.maxhealth
	}
	return p
}

func (p *Player) addHealthWithPointer(health uint) {
	p.health += health
	if p.health > p.maxhealth {
		p.health = p.maxhealth
	}
}

func (p Player) addEnergy(energy uint) Player {
	p.energy += energy
	if p.energy > p.maxenergy {
		p.energy = p.maxenergy
	}
	return p
}

func (p *Player) applyDemage(ammount uint) {
	if p.health-ammount > p.health {
		p.health = 0
	} else {
		p.health -= ammount
	}

	fmt.Println(p.name, "Consume", ammount, "->", p.health)
}

func (p *Player) addEnergyWithPointer(energy uint) {
	p.energy += energy
	if p.energy > p.maxenergy {
		p.energy = p.maxenergy
	}
}

func main() {
	player1 := Player{
		name:      "Ronny",
		health:    100,
		maxhealth: 200,
		energy:    50,
		maxenergy: 150,
	}

	player1PTR := &player1
	fmt.Printf("Current player status:  %+v\n", player1)
	fmt.Printf("Add a Health :  %+v\n", player1.addHealth(25))
	fmt.Println()
	player1PTR.addHealthWithPointer(100)
	fmt.Printf("Add more health with pointer : %+v\n", player1)
	player1PTR.applyDemage(50)
	fmt.Println()
	fmt.Printf("Current Player: %+v\n", player1)
	player1PTR.applyDemage(155)
	fmt.Printf("Current Player: %+v\n", player1)
	fmt.Println()
}
