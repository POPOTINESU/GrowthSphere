package domain

import "github.com/google/uuid"

type User struct {
	userID uuid.UUID
	username string
	accountName string
	
}
