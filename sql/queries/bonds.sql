-- name: InsertBond :one
INSERT INTO bonds (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetBondByID :one
SELECT * FROM bonds WHERE id = $1;

-- name: GetBonds :many
SELECT * FROM bonds ORDER BY created_at;

-- name: GetBondByFields :one
SELECT * FROM bonds WHERE name = $1;

-- name: DeleteBondByID :exec
DELETE FROM bonds WHERE id = $1;