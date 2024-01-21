CREATE TABLE IF NOT EXISTS book_content (
    id SERIAL PRIMARY KEY,
    book_id INT,
    page_num INT,
    content TEXT
);