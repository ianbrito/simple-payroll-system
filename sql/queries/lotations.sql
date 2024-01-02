-- name: InsertLotation :exec
INSERT INTO lotations (id, admission_at, resignation_at, employee_id, department_id, role_id, bond_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetLotationByID :one
SELECT * FROM lotations WHERE id = $1;

-- name: GetLotations :many
SELECT * FROM lotations ORDER BY created_at;

-- name: DeleteLotationByID :exec
DELETE FROM lotations WHERE id = $1;