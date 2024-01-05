package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ianbrito/simple-payroll-system/internal/domain/entity"
	"github.com/ianbrito/simple-payroll-system/internal/infra/db"
)

type DepartmentRepository struct {
	DB      *sql.DB
	Queries db.Queries
}

func NewDepartmentRepository(dbt *sql.DB) *DepartmentRepository {
	return &DepartmentRepository{
		DB:      dbt,
		Queries: *db.New(dbt),
	}
}

func (r *DepartmentRepository) Create(ctx context.Context, department *entity.Department) (*entity.Department, error) {
	record, err := r.Queries.InsertDepartment(
		ctx,
		db.InsertDepartmentParams{
			ID:        department.ID,
			Name:      department.Name,
			Acronym:   department.Acronym,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &entity.Department{
		ID:      record.ID,
		Name:    record.Name,
		Acronym: record.Acronym,
	}, nil
}

func (r *DepartmentRepository) FirstOrCreate(ctx context.Context, department *entity.Department) (*entity.Department, error) {

	record, err := r.Queries.GetDepartmentByFields(ctx, db.GetDepartmentByFieldsParams{
		Name:    department.Name,
		Acronym: department.Acronym,
	})

	if err == sql.ErrNoRows {
		return r.Create(ctx, department)
	}

	if err != nil {
		return nil, err
	}

	return &entity.Department{
		ID:      record.ID,
		Name:    record.Name,
		Acronym: record.Acronym,
	}, nil
}

func (r *DepartmentRepository) FindByID(ctx context.Context, departmentID string) (*entity.Department, error) {
	record, err := r.Queries.GetDepartmentByID(ctx, departmentID)
	if err != nil {
		return nil, err
	}

	department := &entity.Department{
		ID:      record.ID,
		Name:    record.Name,
		Acronym: record.Acronym,
	}

	return department, nil
}

func (r *DepartmentRepository) Find(ctx context.Context) ([]*entity.Department, error) {
	records, err := r.Queries.GetDepartments(ctx)
	if err != nil {
		return nil, err
	}

	var departments []*entity.Department
	for _, department := range records {
		departments = append(departments, &entity.Department{
			ID:      department.ID,
			Name:    department.Name,
			Acronym: department.Acronym,
		})
	}

	return departments, nil
}
