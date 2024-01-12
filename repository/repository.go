package repository

import (
	"context"
	"fmt"
	"myf5/ent"
)

func GetAllPlayers(client *ent.Client) ([]*ent.Player, error) {
	players, err := client.Player.Query().All(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return players, err
}
