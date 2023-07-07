CREATE TABLE IF NOT EXISTS books ON CLUSTER 'permanent-indev' (
       book_id Int64,
       author_id Int64,
       name String,
       page_count Int64
) ENGINE = MergeTree ORDER BY name;

CREATE TABLE IF NOT EXISTS authors ON CLUSTER 'permanent-indev' (
       author_id Int64,
       name String,
       surname String
) ENGINE = MergeTree ORDER BY name;

CREATE TABLE books_distr ON CLUSTER 'permanent-indev' AS books ENGINE = Distributed('permanent-indev', default, books, rand());

CREATE TABLE authors_distr ON CLUSTER 'permanent-indev' AS authors ENGINE = Distributed('permanent-indev', default, authors, rand());
