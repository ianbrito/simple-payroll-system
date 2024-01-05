-- name: InsertRole :one
INSERT INTO roles (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = $1;

-- name: GetRoles :many
SELECT * FROM roles ORDER BY created_at;

-- name: GetRoleByFields :one
SELECT * FROM roles WHERE name = $1;

-- name: DeleteRoleByID :exec
DELETE FROM roles WHERE id = $1;