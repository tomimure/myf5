package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"myf5/ent"
	"myf5/services"
	"net/http"
	"time"
)

func CreatePlayer(w http.ResponseWriter, req *http.Request, client *ent.Client) {
	var requestedData ent.Player
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something bad happened!"))
		return
	}
	err = services.CreatePlayer(client, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func GetAllPlayers(w http.ResponseWriter, req *http.Request, client *ent.Client) {
	players, err := services.GetAllPlayers(client)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something bad happened!"))
		return
	}

	json.NewEncoder(w).Encode(players)

}

func CreateMatch(w http.ResponseWriter, req *http.Request, client *ent.Client) {
	type BodyCreateMatch struct {
		Goalsteam1 int
		Goalsteam2 int
		IdsTeam1   []int
		IdsTeam2   []int
	}

	var requestedData BodyCreateMatch
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Something bad happened!"))
		return
	}
	match := client.Match.Create().SetGoalsTeam1(requestedData.Goalsteam1).SetGoalsTeam2(requestedData.Goalsteam2).SetDate(time.Now()).SaveX(context.Background())
	client.Team.Create().AddPlayerIDs(requestedData.IdsTeam1...).SetMatch(match).SetTeam("1").SaveX(context.Background())
	client.Team.Create().AddPlayerIDs(requestedData.IdsTeam2...).SetMatch(match).SetTeam("2").SaveX(context.Background())

}

func GetPlayersByMatch(w http.ResponseWriter, req *http.Request, client *ent.Client) {
	var requestedData ent.Match
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 Bad Request!"))
		return
	}
	players, err := services.GetPlayersByMatch(client, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(players)

}

func GetMatchedTeams(w http.ResponseWriter, req *http.Request, client *ent.Client) {
	var requestedData services.BodyCreateMatchmacking
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 Bad Request!"))
		return
	}
	teams, err := services.GetMatchmaking(client, requestedData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(teams)

}
