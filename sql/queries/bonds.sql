-- name: InsertBond :exec
INSERT INTO bonds (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4);

-- name: GetBondByID :one
SELECT * FROM bonds WHERE id = $1;

-- name: GetBonds :many
SELECT * FROM bonds ORDER BY created_at;

-- name: DeleteBondByID :exec
DELETE FROM bonds WHERE id = $1;