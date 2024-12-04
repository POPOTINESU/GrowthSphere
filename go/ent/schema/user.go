package schema

import (
	"log"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	MAX_USER_NAME := 20
	MAX_ACCOUNT_NAME := 30

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				uuidV7, err := uuid.NewV7()
				if err != nil {
					log.Fatal("uuid generate error:", err)
				}
				return uuidV7
			}).
			Unique(),
		field.String("username").
			MaxLen(MAX_USER_NAME).
			NotEmpty(),
		field.String("account_name").
			MaxLen(MAX_ACCOUNT_NAME).
			Unique().
			NotEmpty(),
		field.Time("updated_at").
			Default(func() time.Time {
				return time.Now()
			}).
			UpdateDefault(func() time.Time {
				return time.Now()
			}),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
