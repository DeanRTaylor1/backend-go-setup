-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  first_name,
  second_name,
  email
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET 
  username = $2,
  hashed_password = $3,
  first_name = $4,
  second_name = $5,
  email = $6
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
