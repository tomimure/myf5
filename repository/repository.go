package repository

import (
	"context"
	"myf5/ent"
	"myf5/ent/match"
	"myf5/ent/player"
)

func CreatePlayer(client *ent.Client, player *ent.Player) error {
	_, err := client.Player.Create().SetGlobal(player.Global).SetName(player.Name).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func GetAllPlayers(client *ent.Client) ([]*ent.Player, error) {
	players, err := client.Player.Query().All(context.Background())
	if err != nil {
		return nil, err
	}
	return players, err
}

func GetPlayersByMatch(client *ent.Client, matchData *ent.Match) ([]*ent.Player, error) {
	players, err := client.Match.Query().Where(match.ID(matchData.ID)).QueryTeams().QueryPlayers().All(context.Background())
	if err != nil {
		return nil, err
	}
	if len(players) == 0 {
		err = &ent.NotFoundError{}
	}
	return players, err
}

func GetPlayersByID(client *ent.Client, playersIds []int) ([]*ent.Player, error) {
	players, err := client.Player.Query().Where(player.IDIn(playersIds...)).All(context.Background())
	if err != nil {
		return nil, err
	}
	return players, err

}
