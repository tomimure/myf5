package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"myf5/ent"
	"myf5/ent/match"
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
	client.Player.Create().SetGlobal(requestedData.Global).SetName(requestedData.Name).SaveX(context.Background())
	fmt.Println(client.Player.Query().AllX(context.Background()))
}

func listOfPlayers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%v", client.Player.Query().AllX(context.Background()))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
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

func StartServer() {

	http.HandleFunc("/createPlayer", createPlayer)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/playerlist", listOfPlayers)
	http.HandleFunc("/createMatch", createMatch)
	http.HandleFunc("/getplayersbymatch", getPlayersByMatch)

	http.ListenAndServe(":8090", nil)
}
