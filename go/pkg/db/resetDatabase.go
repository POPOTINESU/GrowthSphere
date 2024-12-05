package db

import (
	"GROWTHSPHERE/ent"
	"context"
	"log"
)

// Reset Database for clean test
func ResetDatabase(ctx context.Context, client *ent.Client) {
	_, err := client.User.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to reset database: %v", err)
	}
}
