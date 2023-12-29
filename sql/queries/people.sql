-- name: InsertPerson :exec
INSERT INTO people (id, name, cpf, birth, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetPersonByID :one
SELECT * FROM people WHERE id = $1;

-- name: GetPeople :many
SELECT * FROM people ORDER BY created_at;

-- name: UpdatePerson :exec
UPDATE people SET name = $2, cpf = $3, birth = $4, updated_at = $5 WHERE id = $1;

-- name: DeletePersonByID :exec
DELETE FROM people WHERE id = $1;