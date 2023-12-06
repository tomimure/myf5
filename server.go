package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"myf5/ent"
	"myf5/ent/player"
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
	json.Unmarshal(body, &requestedData)
	client.Player.Create().SetGlobal(requestedData.Global).SetName(requestedData.Name).SaveX(context.Background())
	fmt.Println(client.Player.Query().AllX(context.Background()))
}

func createMatch(w http.ResponseWriter, req *http.Request) {

	type match struct{
		date string
		result string
		players []int
	}

	var requestedData match
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	defer req.Body.Close()
	json.Unmarshal(body, &requestedData)
	client.Match.Create().SetDate(time.Now()).SetResult(requestedData.result).AddPlayers(client.Player.Query().IDs(player)).SaveX(context.Background())

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

func StartServer() {

	http.HandleFunc("/createPlayer", createPlayer)
	http.HandleFunc("/createMatch", createMatch)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/playerlist", listOfPlayers)

	http.ListenAndServe(":8090", nil)
}