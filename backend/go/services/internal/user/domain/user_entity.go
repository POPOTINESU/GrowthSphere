package domain

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	Email       Email
	Password    Password
	Provider    Provider
	Permissions Permission
}