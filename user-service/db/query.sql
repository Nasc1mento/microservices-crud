-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUserById :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :exec
UPDATE users SET
    name = $1,
    email = $2,
    password = $3
WHERE id = $3
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;