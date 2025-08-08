package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Genre struct {
	bun.BaseModel `bun:"table:genres,alias:genres"`

	ID        uuid.UUID `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name      string    `bun:"type:text,notnull"`
	CreatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
}
