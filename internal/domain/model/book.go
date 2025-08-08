package model

import (
	"time"

	"github.com/fitraaditama7/byfood-test/internal/domain/entity"
	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"name"`
	Description string    `json:"description"`
	AuthorID    uuid.UUID `json:"author_id"`
	PublishedAt time.Time `json:"published_at"`
	Picture     *string   `json:"picture"`
	Author      Author    `json:"author"`
	Genres      []Genre   `json:"genres"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) ToEntity() entity.Book {
	return entity.Book{
		ID:          b.ID,
		Title:       b.Title,
		Picture:     b.Picture,
		Published:   b.PublishedAt,
		Description: b.Description,
	}
}

func FromEntityToBook(book entity.Book) Book {
	return Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		AuthorID:    book.AuthorID,
		PublishedAt: book.Published,
		Picture:     book.Picture,
		Author:      FromEntityToAuthor(book.Author),
		Genres:      FromEntitiesToGenres(book.Genres),
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
}
