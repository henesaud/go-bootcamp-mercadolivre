SELECT * FROM movies;

SELECT first_name, last_name, rating from actors;

SELECT title FROM series AS series; 

SELECT first_name, last_name from actors WHERE rating > 7.5;

SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

SELECT title, rating FROM movies ORDER BY rating;

SELECT title FROM movies LIMIT 3;

SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5;

SELECT title, rating FROM movies ORDER BY rating DESC LIMIT 5 OFFSET 5;

SELECT first_name, last_name FROM actors LIMIT 10;

SELECT first_name, last_name FROM actors LIMIT 10 OFFSET 20;

SELECT first_name, last_name FROM actors LIMIT 10 OFFSET 40;

SELECT title, rating FROM movies WHERE title LIKE 'Toy Story';

SELECT first_name, last_name from actors WHERE first_name LIKE 'Sam%';

SELECT title FROM movies WHERE release_date BETWEEN '2004-01-01' AND '2008-12-31';

SELECT title FROM movies WHERE release_date BETWEEN '1988-01-01' AND '2009-12-31' AND
rating > 3 AND awards > 1 ORDER BY rating;

SELECT title FROM movies WHERE release_date BETWEEN '1988-01-01' AND '2009-12-31' AND
rating > 3 AND awards > 1 ORDER BY rating LIMIT 3 OFFSET 10;