package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Book struct {
	bun.BaseModel `bun:"table:books,alias:books"`

	ID          uuid.UUID `bun:"type:uuid,pk,default:uuid_generate_v4()"`
	Title       string    `bun:"type:text,notnull"`
	Description string    `bun:"type:text, notnull"`
	AuthorID    uuid.UUID `bun:"type:uuid,notnull"`
	Published   time.Time `bun:"type:date,notnull"`
	Picture     *string   `bun:"type:text"`
	CreatedAt   time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"type:timestamptz,notnull,default:current_timestamp"`

	Author Author  `bun:"rel:belongs-to,join:author_id=id"`
	Genres []Genre `bun:"m2m:books_genres,join:book_id=id,join:genre_id=id"`
}
