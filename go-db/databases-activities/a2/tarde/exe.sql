USE movies_db;

INSERT INTO movies
(title, rating, awards, release_date, genre_id) VALUES
('Taxi driver', 4, 3, '1999-01-01', 3);

INSERT INTO genres (name, ranking) VALUES
('Romance', 13);

UPDATE movies SET genre_id = 13 WHERE id = 22;

UPDATE actors SET favorite_movie_id = 22 WHERE id = 3;

CREATE TEMPORARY TABLE movies_table SELECT * FROM movies;

DELETE FROM movies_table WHERE awards < 5;

SELECT * FROM movies_table;

SELECT genres.name FROM genres WHERE (SELECT count(movies.id) FROM movies WHERE movies.genre_id = genres.id) > 0;

SELECT actors.first_name, movies.title AS favorite_movie, movies.awards FROM actors
INNER JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

EXPLAIN SELECT * FROM movies;

EXPLAIN DELETE FROM movies_table WHERE awards < 5;

EXPLAIN SELECT genres.name FROM genres WHERE (SELECT count(movies.id) FROM movies WHERE movies.genre_id = genres.id) > 0;

EXPLAIN SELECT actors.first_name, movies.title AS favorite_movie, movies.awards FROM actors
INNER JOIN movies ON actors.favorite_movie_id = movies.id
WHERE movies.awards > 3;

ALTER TABLE movies ADD INDEX index_name(title);

SHOW INDEX FROM movies;