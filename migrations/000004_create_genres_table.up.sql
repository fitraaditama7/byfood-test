CREATE TABLE IF NOT EXISTS genres (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TRIGGER trigger_set_updated_at
    BEFORE UPDATE ON genres
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at_timestamp();

INSERT INTO genres (name) VALUES
  ('Science Fiction'),
  ('Fantasy'),
  ('Mystery'),
  ('Thriller'),
  ('Romance'),
  ('Horror'),
  ('Historical Fiction'),
  ('Biography'),
  ('Non-Fiction'),
  ('Self-Help'),
  ('Adventure'),
  ('Young Adult'),
  ('Children’s'),
  ('Dystopian'),
  ('Crime'),
  ('Graphic Novel'),
  ('Memoir'),
  ('Poetry'),
  ('Classic'),
  ('Drama'),
  ('Satire'),
  ('Paranormal'),
  ('Cookbook'),
  ('Travel'),
  ('Religion / Spirituality'),
  ('Science'),
  ('Philosophy');