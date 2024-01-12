package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"myf5/ent"
	"myf5/ent/match"
	"myf5/ent/player"
	"myf5/services"
	"net/http"
	"time"
)

func createPlayer(w http.ResponseWriter, req *http.Request) {

	var requestedData ent.Player
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	err = json.Unmarshal(body, &requestedData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Something bad happened!"))
		return
	}
	services.NewPlayerService(client).Create(&requestedData)

}

func createMatch(w http.ResponseWriter, req *http.Request) {
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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 - Something bad happened!"))
		return
	}
	match := client.Match.Create().SetGoalsTeam1(requestedData.Goalsteam1).SetGoalsTeam2(requestedData.Goalsteam2).SetDate(time.Now()).SaveX(context.Background())
	client.Team.Create().AddPlayerIDs(requestedData.IdsTeam1...).SetMatch(match).SetTeam("1").SaveX(context.Background())
	client.Team.Create().AddPlayerIDs(requestedData.IdsTeam2...).SetMatch(match).SetTeam("2").SaveX(context.Background())

}

func getPlayersByMatch(w http.ResponseWriter, req *http.Request) {
	var requestedData ent.Match
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
	match := client.Match.Query().Where(match.ID(requestedData.ID)).QueryTeams().QueryPlayers().AllX(context.Background())

	json.NewEncoder(w).Encode(match)

}

func createTeamsMatchs(w http.ResponseWriter, req *http.Request) {
	type bodyCreateMatchmacking struct {
		PlayersByTeam int
		PlayersIds    []int
	}
	var requestedData bodyCreateMatchmacking
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
	if len(requestedData.PlayersIds) == (2 * requestedData.PlayersByTeam) {
		json.NewEncoder(w).Encode(Matchmaking(client.Player.Query().Where(player.IDIn(requestedData.PlayersIds...)).AllX(context.Background())))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("400 Bad Request, incorrect number of ids"))
		return
	}

}

func StartServer() {

	http.HandleFunc("/createPlayer", createPlayer)
	http.HandleFunc("/createMatch", createMatch)
	http.HandleFunc("/getplayersbymatch", getPlayersByMatch)
	http.HandleFunc("/createMatchmacking", createTeamsMatchs)

	http.ListenAndServe(":8090", nil)
	defer client.Close()
}
