ALTER TABLE book_content 
  ADD CONSTRAINT fk_book
  FOREIGN KEY (book_id) 
  REFERENCES books(id);