-- name: InsertRole :exec
INSERT INTO roles (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4);

-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = $1;

-- name: GetRoles :many
SELECT * FROM roles ORDER BY created_at;

-- name: DeleteRoleByID :exec
DELETE FROM roles WHERE id = $1;