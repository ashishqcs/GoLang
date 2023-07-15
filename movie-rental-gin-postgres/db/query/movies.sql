-- name: CreateMovie :one
INSERT INTO movies (
  id, title, released, genre, actors, year, price, quantity
) VALUES(
  $1, $2, $3, $4, $5, $6, $7, $8
) 
ON CONFLICT (id) 
DO 
   UPDATE SET title = $2, released = $3, genre = $4, actors = $5, year = $6, price = $7, quantity = $8
RETURNING *;

-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
WHERE (CASE WHEN @lk_genre::bool THEN genre LIKE '%' || @genre || '%' ELSE TRUE END)
  AND (CASE WHEN @lk_actors::bool THEN actors LIKE '%' || @actor || '%' ELSE TRUE END)
  AND (CASE WHEN @eq_year::bool THEN year = @year ELSE TRUE END)
LIMIT $1
OFFSET $2;