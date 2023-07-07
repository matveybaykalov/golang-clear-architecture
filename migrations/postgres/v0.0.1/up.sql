CREATE TABLE IF NOT EXISTS books (
       book_id INT,
       author_id INT,
       name TEXT,
       page_count INT
);

CREATE TABLE IF NOT EXISTS authors (
       author_id INT,
       name TEXT,
       surname TEXT
);
