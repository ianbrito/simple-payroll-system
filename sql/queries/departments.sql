-- name: InsertDepartment :exec
INSERT INTO departments (id, name, acronym, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);

-- name: GetDepartmentByID :one
SELECT * FROM departments WHERE id = $1;

-- name: GetDepartments :many
SELECT * FROM departments ORDER BY created_at;

-- name: DeleteDepartmentByID :exec
DELETE FROM departments WHERE id = $1;