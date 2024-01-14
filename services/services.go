package services

import (
	"errors"
	"myf5/ent"
	"myf5/matchmaking"
	"myf5/repository"
)

func CreatePlayer(client *ent.Client, player *ent.Player) error {
	err := repository.CreatePlayer(client, player)
	if err != nil {
		return err
	}
	return nil

}

func GetAllPlayers(client *ent.Client) ([]*ent.Player, error) {
	players, err := repository.GetAllPlayers(client)
	return players, err
}

func GetPlayersByMatch(client *ent.Client, matchData *ent.Match) ([]*ent.Player, error) {
	players, err := repository.GetPlayersByMatch(client, matchData)
	return players, err
}

type BodyCreateMatchmacking struct {
	PlayersByTeam int
	PlayersIds    []int
}

func GetMatchmaking(client *ent.Client, data BodyCreateMatchmacking) ([][]*ent.Player, error) {
	//TODO: Fix case of sending a non existing id
	if len(data.PlayersIds) == (2 * data.PlayersByTeam) {
		players, err := repository.GetPlayersByID(client, data.PlayersIds)
		if err != nil {
			return nil, err
		}
		teams := matchmaking.Matchmaking(players)
		return teams, err
	}
	err := errors.New("incorrect number of ids")
	return nil, err
}
