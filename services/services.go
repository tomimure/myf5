package services

import (
	"context"
	"fmt"
	"myf5/ent"
)

type PlayerService struct {
	client *ent.Client
}

func NewPlayerService(client *ent.Client) *PlayerService {
	return &PlayerService{
		client: client,
	}
}

func (s *PlayerService) Create(player *ent.Player) error {
	fmt.Println(player)
	_, err := s.client.Player.Create().SetGlobal(player.Global).SetName(player.Name).Save(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
