package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Author struct {
	bun.BaseModel `bun:"table:authors,alias:authors"`

	ID        uuid.UUID `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Name      string    `bun:"type:text,notnull"`
	Picture   *string   `bun:"type:text"`
	CreatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
}
