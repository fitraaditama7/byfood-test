package model

import (
	"time"

	"github.com/fitraaditama7/byfood-test/internal/domain/entity"
	"github.com/google/uuid"
)

type Genre struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (g *Genre) ToEntity() entity.Genre {
	return entity.Genre{
		ID:        g.ID,
		Name:      g.Name,
		CreatedAt: g.CreatedAt,
		UpdatedAt: g.UpdatedAt,
	}
}

func FromEntityToGenre(genre entity.Genre) Genre {
	return Genre{
		ID:        genre.ID,
		Name:      genre.Name,
		CreatedAt: genre.CreatedAt,
		UpdatedAt: genre.UpdatedAt,
	}
}

func FromEntitiesToGenres(entities []entity.Genre) []Genre {
	result := make([]Genre, len(entities))
	for i, g := range entities {
		result[i] = FromEntityToGenre(g)
	}
	return result
}
