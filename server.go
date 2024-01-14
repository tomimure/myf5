package main

import (
	"fmt"
	"myf5/controller"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/createplayer", func(w http.ResponseWriter, r *http.Request) {
		controller.CreatePlayer(w, r, client)
	})
	http.HandleFunc("/getallplayers", func(w http.ResponseWriter, r *http.Request) {
		controller.GetAllPlayers(w, r, client)
	})
	http.HandleFunc("/getplayersbymatch", func(w http.ResponseWriter, r *http.Request) {
		controller.GetPlayersByMatch(w, r, client)
	})
	http.HandleFunc("/getmatchedteams", func(w http.ResponseWriter, r *http.Request) {
		controller.GetMatchedTeams(w, r, client)
	})
	// http.HandleFunc("/createMatch", controller.CreateMatch)
	fmt.Println("Server 8090 listen")
	http.ListenAndServe(":8090", nil)
	defer client.Close()
}
