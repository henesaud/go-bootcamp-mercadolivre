SELECT title AS movie_title, name AS genre_name
FROM movies
INNER JOIN genres ON
movies.genre_id = genres.id;

SELECT episodes.title AS episode_title, 
CONCAT(actors.first_name, ' ', actors.last_name) AS actor_name
FROM actors
INNER JOIN actor_episode ON actors.id = actor_episode.actor_id
INNER JOIN episodes ON actor_episode.episode_id = episodes.id;

SELECT series.title AS serie_title, count(episodes.id) AS episodes_quantity
FROM episodes
INNER JOIN seasons ON episodes.season_id = seasons.id
INNER JOIN series ON seasons.serie_id = series.id
GROUP BY series.id;

SELECT genres.name, COUNT(movies.id) movies_quantity
FROM genres
INNER JOIN movies ON genres.id = movies.genre_id
GROUP BY genres.id
HAVING movies_quantity >= 3;

SELECT DISTINCT CONCAT(actors.first_name, ' ', actors.last_name) AS actor_name
FROM actors
INNER JOIN actor_movie ON actor_movie.actor_id = actors.id
INNER JOIN movies ON movies.id = actor_movie.movie_id
WHERE movies.title LIKE '%Star Wars%';