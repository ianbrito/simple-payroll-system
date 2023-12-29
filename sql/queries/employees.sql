-- name: InsertEmployee :exec
INSERT INTO employees (id, enrollment, person_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5);

-- name: GetEmployeeByID :one
SELECT * FROM employees WHERE id = $1;

-- name: GetEmployees :many
SELECT * FROM employees ORDER BY created_at;

-- name: DeleteEmployeeByID :exec
DELETE FROM employees WHERE id = $1;