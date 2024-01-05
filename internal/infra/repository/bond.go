package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/ianbrito/simple-payroll-system/internal/domain/entity"
	"github.com/ianbrito/simple-payroll-system/internal/infra/db"
)

type BondRepository struct {
	DB      *sql.DB
	Queries db.Queries
}

func NewBondRepository(dbt *sql.DB) *BondRepository {
	return &BondRepository{
		DB:      dbt,
		Queries: *db.New(dbt),
	}
}

func (r *BondRepository) FirstOrCreate(ctx context.Context, bond *entity.Bond) (*entity.Bond, error) {

	record, err := r.Queries.GetBondByFields(ctx, bond.Name)

	if err == sql.ErrNoRows {
		return r.Create(ctx, bond)
	}

	if err != nil {
		return nil, err
	}

	return &entity.Bond{ID: record.ID, Name: record.Name}, nil
}

func (r *BondRepository) Create(ctx context.Context, bond *entity.Bond) (*entity.Bond, error) {
	record, err := r.Queries.InsertBond(
		ctx,
		db.InsertBondParams{
			ID:        bond.ID,
			Name:      bond.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	)

	if err != nil {
		return nil, err
	}

	return &entity.Bond{
		ID:   record.ID,
		Name: record.Name,
	}, nil
}

func (r *BondRepository) FindByID(ctx context.Context, bondID string) (*entity.Bond, error) {
	record, err := r.Queries.GetBondByID(ctx, bondID)

	if err != nil {
		return nil, err
	}

	bond := &entity.Bond{
		ID:   record.ID,
		Name: record.Name,
	}

	return bond, nil
}

func (r *BondRepository) All(ctx context.Context) ([]*entity.Bond, error) {
	records, err := r.Queries.GetBonds(ctx)
	if err != nil {
		return nil, err
	}

	var bonds []*entity.Bond
	for _, b := range records {
		bonds = append(bonds, &entity.Bond{
			ID:   b.ID,
			Name: b.Name,
		})
	}

	return bonds, nil
}
