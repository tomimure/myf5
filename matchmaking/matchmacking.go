package matchmaking

import (
	"myf5/ent"
	"sort"
)

func Matchmaking(players []*ent.Player) [][]*ent.Player {
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
	var response [][]*ent.Player
	response = append(response, team1, team2)
	return response

}
func getGlobal(team []*ent.Player) int {
	var sum int
	for i := 0; i < len(team); i++ {
		sum = sum + team[i].Global
	}
	return sum
}
