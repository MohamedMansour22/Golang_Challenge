package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Transaction struct {
	bun.BaseModel `bun:"table:transactions"`
	ID            uuid.UUID `bun:"id,notnull,pk,type:uuid,default:gen_random_uuid()"`
	Amount        float64   `bun:"amount" `
	Currency      string    `bun:"currency" `
	CreatedAt     string    `bun:"createdat" `
	Status        bool      `bun:"IsActive" `
}
