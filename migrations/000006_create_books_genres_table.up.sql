CREATE TABLE IF NOT EXISTS books_genres (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    book_id UUID NOT NULL,
    genre_id UUID NOT NULL,
    CONSTRAINT fk_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    CONSTRAINT fk_genre FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
);
