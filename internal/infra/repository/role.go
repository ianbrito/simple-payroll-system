package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ianbrito/simple-payroll-system/internal/domain/entity"
	"github.com/ianbrito/simple-payroll-system/internal/infra/db"
)

type RoleRepository struct {
	DB      *sql.DB
	Queries db.Queries
}

func NewRoleRepository(dbt *sql.DB) *RoleRepository {
	return &RoleRepository{
		DB:      dbt,
		Queries: *db.New(dbt),
	}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	record, err := r.Queries.InsertRole(
		ctx,
		db.InsertRoleParams{
			ID:        role.ID,
			Name:      role.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)
	if err != nil {
		return nil, err
	}
	return &entity.Role{ID: record.ID, Name: record.Name}, nil
}

func (r *RoleRepository) FirstOrCreate(ctx context.Context, role *entity.Role) (*entity.Role, error) {
	record, err := r.Queries.GetRoleByFields(ctx, role.Name)

	if err == sql.ErrNoRows {
		return r.CreateRole(ctx, role)
	}

	if err != nil {
		return nil, err
	}

	return &entity.Role{ID: record.ID, Name: record.Name}, nil
}

func (r *RoleRepository) FindRoleByID(ctx context.Context, roleID string) (*entity.Role, error) {
	res, err := r.Queries.GetRoleByID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	role := &entity.Role{
		ID:   res.ID,
		Name: res.Name,
	}

	return role, nil
}

func (r *RoleRepository) All(ctx context.Context) ([]*entity.Role, error) {
	records, err := r.Queries.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	var roles []*entity.Role
	for _, role := range records {
		roles = append(roles, &entity.Role{
			ID:   role.ID,
			Name: role.Name,
		})
	}

	return roles, nil
}
