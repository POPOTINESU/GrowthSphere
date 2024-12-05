package db

import (
	"GROWTHSPHERE/ent"
	"GROWTHSPHERE/pkg/config"
	"fmt"
	"log"
)

func CreateTestClient() *ent.Client {
	client, err := ent.Open("postgres", readTestDatabaseEnv())
	if err != nil {
		log.Fatalf("failed opening test database: %v", err)
	}
	return client
}

func readTestDatabaseEnv() string {
	host := config.GetEnv("POSTGRES_HOST")
	port := config.GetEnv("POSTGRES_DB")
	user := config.GetEnv("POSTGRES_USER")
	dbname := config.GetEnv("POSTGRES_DB")
	password := config.GetEnv("POSTGRES_PASSWORD")

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)
}
