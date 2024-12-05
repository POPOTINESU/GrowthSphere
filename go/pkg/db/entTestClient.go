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
	host := config.GetEnv("TEST_DB_HOST")
	port := config.GetEnv("TEST_DB_PORT")
	user := config.GetEnv("TEST_DB_USER")
	dbname := config.GetEnv("TEST_DB_NAME")
	password := config.GetEnv("TEST_DB_PASSWORD")

	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password,
	)
}
