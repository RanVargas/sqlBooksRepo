CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS book_content (
    id SERIAL PRIMARY KEY,
    book_id INT,
    page_num INT,
    content TEXT
);
ALTER TABLE book_content 
  ADD CONSTRAINT fk_book
  FOREIGN KEY (book_id) 
  REFERENCES books(id);