package db

import (
	ent "GROWTHSPHERE/ent"
	"GROWTHSPHERE/pkg/config"
	"fmt"
	"log"
)

func CreateClient() ent.Client {
	client, err := ent.Open("postgres", readDatabaseEnv())
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return *client
}

func readDatabaseEnv() string {
	host := config.GetEnv("DB_HOST")
	port := config.GetEnv("DB_PORT")
	user := config.GetEnv("DB_USER")
	dbname := config.GetEnv("DB_NAME")
	password := config.GetEnv("DB_PASSWORD")

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)
}
