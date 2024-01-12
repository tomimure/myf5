package services

import (
	"context"
	"fmt"
	"myf5/ent"
	"myf5/repository"
)

func CreatePlayer(client *ent.Client, player *ent.Player) error {
	_, err := client.Player.Create().SetGlobal(player.Global).SetName(player.Name).Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func GetAllPlayers(client *ent.Client) ([]*ent.Player, error) {
	players, err := repository.GetAllPlayers(client)
	return players, err
}
