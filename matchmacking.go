package main

import (
	"fmt"
	"myf5/ent"
	"sort"
)

func Matchmaking(players []*ent.Player) {
	fmt.Println(players)
	var team1 []*ent.Player
	var team2 []*ent.Player
	sort.Slice(players, func(i, j int) bool {
		return players[i].Global > players[j].Global
	})
	for i := 0; i < len(players); i++ {
		if getGlobal(team1) > getGlobal(team2) {
			team2 = append(team2, players[i])
		} else {
			team1 = append(team1, players[i])
		}
	}
	fmt.Println(players)
	fmt.Println(team1)
	fmt.Println(team2)

}
func getGlobal(team []*ent.Player) int {
	var sum int
	for i := 0; i < len(team); i++ {
		sum = sum + team[i].Global
	}
	return sum
}
