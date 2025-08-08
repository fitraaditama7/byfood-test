package model

import (
	"time"

	"github.com/fitraaditama7/byfood-test/internal/domain/entity"
	"github.com/google/uuid"
)

type Author struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Picture   *string   `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Author) ToEntity() entity.Author {
	return entity.Author{
		ID:        a.ID,
		Name:      a.Name,
		Picture:   a.Picture,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func FromEntityToAuthor(author entity.Author) Author {
	return Author{
		ID:        author.ID,
		Name:      author.Name,
		Picture:   author.Picture,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}
}
