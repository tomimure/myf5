package main

import (
	"context"
	"fmt"
	"log"
	"myf5/ent"

	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func main() {
	// TODO: levantar variables de entorno
	setupDb()
	// Matchmaking(client.Player.Query().AllX(context.Background()))
	StartServer()

}

func setupDb() error {
	var err error
	client, err = ent.Open("mysql", DatabaseDSN("root", "changeme", "127.0.0.1", "4306", "myf5"))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
		return err
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
		return err
	}
	return nil
}

func DatabaseDSN(user, pass, host, port, name string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?timeout=10s&charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		host,
		port,
		name,
	)
}
