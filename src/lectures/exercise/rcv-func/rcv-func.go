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
	Health, MaxHealth, Energy, MaxEnergy int
	Name                                 string
}

//func helthCh(i int, j int, player *Player) {
func (player *Player) maxStatSet(i int, j int) {
	player.MaxEnergy = i
	player.MaxHealth = j

}
func (player *Player) statCH(i int, j int) {
	fmt.Println(player)
	player.Health += i
	player.Energy += j
	if player.Health > player.MaxHealth {
		player.Health = player.MaxHealth
	}
	if player.Health < 0 {
		player.Health = 0
	}

	if player.Energy > player.MaxEnergy {
		player.Energy = player.MaxEnergy
	}
	if player.Energy < 0 {
		player.Energy = 0
	}

	fmt.Println(player)
}

func main() {
	player1 := Player{}
	player1.maxStatSet(5, 6)
	player1.statCH(3, -1)
	fmt.Println(player1)
}
